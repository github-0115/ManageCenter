/**
 * Created by huang on 2016/8/29.
 */
$(document).ready(function () {
    var defaults = {
        viewUrl: "/demo/detail",
        page: 1,
        rows: 100,
        prev_text: "前一页",
        next_text: "下一页",
        num_edge_entries: 2, //边缘页数
        num_display_entries: 5, //主体页数
        items_per_page: 100,
    };
    Date.prototype.format = function (fmt) {
        var o = {
            "M+": this.getMonth() + 1,                 //月份
            "d+": this.getDate(),                    //日
            "h+": this.getHours(),                   //小时
            "m+": this.getMinutes(),                 //分
            "s+": this.getSeconds(),                 //秒
            "q+": Math.floor((this.getMonth() + 3) / 3), //季度
            "S": this.getMilliseconds()             //毫秒
        };
        if (/(y+)/.test(fmt))
            fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
        for (var k in o)
            if (new RegExp("(" + k + ")").test(fmt))
                fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
        return fmt;
    }

    var time = "";
    var isFirstLoad = false
    var isScreen = false
    var sord = "desc"
    var currentStatus="identify"
    var mydate = new Date();
    time = mydate.format("yyyy-MM-dd"); 
     var resultClass = {
        "p": "result-bg-red",
        "s": "result-bg-yellow",
        "n": "result-bg-green",
    }
     var cnResult = {
        "p": "色情(P)",
        "s": "性感(S)",
        "n": "正常(N)",
    }
    function initImg(data) {
        var div = $('#img');
        div.empty()
        data.result.forEach(function (item, idx) {
            var img = "";
            img += '<div class="box" >'
            img += '<div class="info">'
            img += '<label class="label infolabel ">'
            img += '<span class="' + resultClass[item.res] + '">' + cnResult[item.res] + '</span>'
            if (item.review)
            img += '<span class="fa fa-exclamation-circle result-bg-danger"></span>'
            img += '</label>'
            img += '<label class="labelrate ">'
            img += '<span class="rate">' + (item.rate * 100).toFixed(2) + '%</span>'
            img += '</label>'
            img += '<label class="datetime ">'
            img += '<span class="time">' + item.created_at+ '</span>'
            img += '</label>'
            img += '</div>'
            img += '<div class="imgbox">'
            img += '<img class="imgview" alt="..." src=' + item.url + '   />'
            img += '</div>'
            img += '</div>'
            div.append(img)
        })
    }
      function initFacesImg(data) {
        var div = $('#img');
        div.empty()
        data.result.forEach(function (item, idx) {
            var img = "";
            img += '<div class="box" >'
            img += '<div class="faceinfo">'
            img += '<label class="datetime ">'
            img += '<span class="time">' + item.created_at+ '</span>'
            img += '</label>'
            img += '</div>'
            img += '<div class="imgbox" id="img'+idx+'">'
            img += '</div>'
            img += '</div>'
            div.append(img)
            var image = new Image();
            image.src = item.url;
            image.onload = function () {
                drawImage($("#img"+idx),image,item.face_bounds)
            }
            
        })
    }

     function str_format(template, obj) {
        return template.replace(/#\{([^\{\}]+)\}/g, function(match, part) {
            return obj[part];
        });
    }

    function drawImage(container,image,face_bounds){
     container.empty().append(str_format('<canvas width="#{width}" height="#{height}"></canvas>',
                {width: container.width(), height: container.height()}));

            var canvas = container.find("canvas")[0];
            canvas.width = container.width();
            canvas.height = container.height();
            var context = canvas.getContext("2d");
            var xOffset=0
            var yOffset=0
            var scale=1
            if (image.width > container.width() || image.height > container.height()) {
                scale = Math.min(container.width() * 1.0 / image.width, container.height() * 1.0 / image.height);
                context.scale(scale, scale);
            }
            xOffset=(container.width()-image.width*scale)/2/scale
            yOffset=(container.height()-image.height*scale)/2/scale

            context.drawImage(image, xOffset, yOffset);
            if (currentStatus=="faces_landmarks"){
                 face_bounds.forEach(function(item,idx){
                    item.landmarks.forEach(function(item,idx){
                    
                        var x=item.x*image.width
                        var y=item.y* image.height;
                        context.beginPath();
                        context.arc(x+xOffset, y+yOffset, 3.5, 0, Math.PI * 2, true);
                        context.closePath();
                        context.fillStyle = '#00b4ff';
                        context.fill();
                    
                    })
                 })
            }else{
            face_bounds.forEach(function(item,idx){
                var x=item.lefttop.x*image.width
                var y=item.lefttop.y* image.height;
                var xr=item.rightbottom.x*image.width
                var yr=item.rightbottom.y* image.height;
                context.strokeStyle = "#00b4ff";
                context.lineWidth = 3;
                context.strokeRect(x+xOffset,y+yOffset,xr-x,yr-y);
            })
        }
    }

    function pageselectCallback(page_index, jq) {
            if (isFirstLoad || page_index != 0) {
                $.ajax({
                    type: 'get',
                    url: defaults.viewUrl,
                    data: {
                        "time": time,
                        "sord":sord,
                        "type":currentStatus,
                        "page": page_index + 1,
                        "rows": defaults.rows,
                    },
                    headers: {
                        "LoginToken": loginToken,
                        "Access-Control-Allow-Headers": "LoginToken",
                    },
                    beforeSend: function (data) {
                        $.LoadingOverlay("show");
                    },
                    success: function (data) {
                        if(currentStatus=="identify")
                            initImg(data)
                        else
                            initFacesImg(data)
                    },
                    complete: function () {
                        $.LoadingOverlay("hide");
                    },
                    error: function (data) {
                    },
                });
                isFirstLoad = true
            }
        return false;
    }

    function initPagination(data) {
        $("#Pagination").pagination(data.records, {
            prev_text: defaults.prev_text,
            next_text: defaults.next_text,
            num_edge_entries: defaults.num_edge_entries, //边缘页数
            num_display_entries: defaults.num_display_entries, //主体页数
            callback: pageselectCallback,
            items_per_page: defaults.items_per_page //每页显示1项
        });
    }
    initPhoto()
function initPhoto(){
    $.ajax({
        type: 'get',
        url: defaults.viewUrl,
        data: {
            "time": time,
            "sord":sord,
            "type":currentStatus,
            "page": defaults.page,
            "rows": defaults.rows,
        },
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        beforeSend: function () {
            $.LoadingOverlay("show");
        },
        success: function (data) {
            $("#datetimepicker").attr("value", time)
            $("#"+currentStatus).addClass("button-select");
            initPagination(data)
            initImg(data)
        },
        complete: function () {
            $.LoadingOverlay("hide");
        },
        error: function (data) {
        },
    });
}   
$("#identify").click(function () {
    currentStatus="identify"
    $(".reviewcheck").hide();
    $("#identify").addClass("button-select");
    $("#faces").removeClass("button-select");
        $.ajax({
            type: 'get',
            url: defaults.viewUrl,
            data: {
                "time": time,
                "type":currentStatus,
                "sord":sord,
                "page": defaults.page,
                "rows": defaults.rows,
            },
            headers: {
                "LoginToken": loginToken,
                "Access-Control-Allow-Headers": "LoginToken",
            },
            beforeSend: function () {
                $.LoadingOverlay("show");
            },
            success: function (data) {
                $("#timeup").hide();
                $("#timedown").show();
                initPagination(data)
                initImg(data)
            },
            complete: function () {
                $.LoadingOverlay("hide");
            },
            error: function (data) {
            },
        });
    });
$("#faces").click(function () {
    $("#faces").addClass("button-select");
    $("#identify").removeClass("button-select");
    currentStatus="faces"
    $(".reviewcheck").show();
    $(".showreview").prop('checked', false);
    $(".checktype").prop('checked', true);
    initCheckBox(".showreview")
    initCheckBox(".checktype")
        $.ajax({
            type: 'get',
            url: defaults.viewUrl,
            data: {
                "time": time,
                "type":currentStatus,
                "sord":sord,
                "page": defaults.page,
                "rows": defaults.rows,
            },
            headers: {
                "LoginToken": loginToken,
                "Access-Control-Allow-Headers": "LoginToken",
            },
            beforeSend: function () {
                $.LoadingOverlay("show");
            },
            success: function (data) {
                $("#timeup").hide();
                $("#timedown").show();
                initPagination(data)
                initFacesImg(data)
            },
            complete: function () {
                $.LoadingOverlay("hide");
            },
            error: function (data) {
            },
        });
    });
    $("#timeup").click(function () {
        sord = "desc"
        $.ajax({
            type: 'get',
            url: defaults.viewUrl,
            data: {
                "time": time,
                "sord":sord,
                "type":currentStatus,
                "page": defaults.page,
                "rows": defaults.rows,
            },
            headers: {
                "LoginToken": loginToken,
                "Access-Control-Allow-Headers": "LoginToken",
            },
            beforeSend: function () {
                $.LoadingOverlay("show");
            },
            success: function (data) {
                $("#timeup").hide();
                $("#timedown").show();
                initPagination(data)
                if(currentStatus=="identify")
                    initImg(data)
                else
                    initFacesImg(data)
            },
            complete: function () {
                $.LoadingOverlay("hide");
            },
            error: function (data) {
            },
        });
    });

    $("#timedown").click(function () {
        sord = "asc"
        $.ajax({
            type: 'get',
            url: defaults.viewUrl,
            data: {
                "time": time,
                "sord":sord,
                "type":currentStatus,
                "page": defaults.page,
                "rows": defaults.rows,
            },
            headers: {
                "LoginToken": loginToken,
                "Access-Control-Allow-Headers": "LoginToken",
            },
            beforeSend: function () {
                $.LoadingOverlay("show");
            },
            success: function (data) {
                $("#timedown").hide();
                $("#timeup").show();
                initPagination(data)
                if(currentStatus=="identify")
                    initImg(data)
                else
                    initFacesImg(data)
            },
            complete: function () {
                $.LoadingOverlay("hide");
            },
            error: function (data) {
            },
        });
    });
    function initCheckBox(dom) {
        if ($(dom).is(':checked')) {
            $(dom).prop("disabled", true);
        } else {
            $(dom).prop("disabled", false);
        }
    }
    $(".showreview").change(function () {
        currentStatus="faces_landmarks"
        isFirstLoad = false
        if ($(".showreview").is(':checked')) {
            $(".checktype").attr("checked", false);
            $(".showreview").prop("disabled", true);
            $(".checktype").prop("disabled", false);
        }
        $.ajax({
            type: 'get',
            url: defaults.viewUrl,
            data: {
                "time": time,
                "sord":sord,
                "type":currentStatus,
                "page": defaults.page,
                "rows": defaults.rows,
            },
            headers: {
                "LoginToken": loginToken,
                "Access-Control-Allow-Headers": "LoginToken",
            },
            beforeSend: function () {
                $.LoadingOverlay("show");
            },
            success: function (data) {
                initPagination(data)
                initFacesImg(data)
            },
            complete: function () {
                $.LoadingOverlay("hide");
            },
            error: function (data) {
            },

        });
    })

    $(".checktype").change(function () {
        currentStatus="faces"
        isFirstLoad = false
        if ($(".checktype").is(':checked')) {
            $(".showreview").attr("checked", false);
            $(".checktype").prop("disabled", true);
            $(".showreview").prop("disabled", false);
            review = false
        }
        $.ajax({
            type: 'get',
            url: defaults.viewUrl,
            data: {
               "time": time,
                "sord":sord,
                "type":currentStatus,
                "page": defaults.page,
                "rows": defaults.rows,
            },
            headers: {
                "LoginToken": loginToken,
                "Access-Control-Allow-Headers": "LoginToken",
            },
            beforeSend: function () {
                $.LoadingOverlay("show");
            },
            success: function (data) {
                initPagination(data)
                initFacesImg(data)
            },
            complete: function () {
                $.LoadingOverlay("hide");
            },
            error: function (data) {
            },

        });

    })
    $('#datetimepicker').datetimepicker({
        format: 'Y-m-d',
        timepicker: false,
        onSelectDate: function (dp, $input) {
            time = $input.val();
            $.ajax({
                type: 'get',
                url: defaults.viewUrl,
                data: {
                    "time": time,
                    "sord": sord,
                    "type":currentStatus,
                    "page": defaults.page,
                    "rows": defaults.rows,
                },
                headers: {
                    "LoginToken": loginToken,
                    "Access-Control-Allow-Headers": "LoginToken",
                },
                beforeSend: function () {
                    $.LoadingOverlay("show");
                },
                success: function (data) {
                    initPagination(data)
                    if(currentStatus=="identify")
                        initImg(data)
                    else
                        initFacesImg(data)
                },
                complete: function () {
                    $.LoadingOverlay("hide");
                },
                error: function (data) {
                },

            });
        }
    });

});
