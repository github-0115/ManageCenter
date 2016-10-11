/**
 * Created by huang on 2016/7/25.
 */
$(document).ready(function () {
    var defaults = {
        viewUrl: "/api/porn/detail",
        screenUrl: "/api/porn/screen_detail",
        searchUrl: "/stat_search_users",
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

    var type = "";
    var time = "";
    var review = false;
    var _review = "";
    var isFirstLoad = false
    var isScreen = false
    var pBoth = 1
    var sBoth = 1
    var isAll = false
    var sort = ""
    var sord = "desc"
    var cnResult = {
        "p": "色情(P)",
        "s": "性感(S)",
        "isScreen": "拍屏(isScreen)",
    }
    var resultClass = {
        "p": "result-bg-red",
        "s": "result-bg-yellow",
        "isScreen": "result-bg-purple",
    }
    if (getUrlParam("username") == null) {
        username ="yk"
    } else {
        username = getUrlParam("username");
    }
    var mydate = new Date();
    if (getUrlParam("time") == null) {
        time = mydate.format("yyyy-MM-dd");
    } else {
        time = getUrlParam("time");
    }
    if (getUrlParam("type") == null) {
        type = "p"
    } else {
        type = getUrlParam("type");
    }
    function getUrlParam(name) {
        var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
        var r = window.location.search.substr(1).match(reg);  //匹配目标参数
        if (r != null) return unescape(r[2]);
        return null; //返回参数值
    }

    function initImg(data) {
        var div = $('#img');
        div.empty()
        /*  data.result=[{
         "res":"s",
         "review":true,
         "p":0.009381188,
         "s":0.9823659,
         "n":0.00825284,
         "url":"http://statsimgs.oss-cn-hangzhou.aliyuncs.com/cb58ed4b4596bcbb87dfcae18069b4de",
         "created_at":111111111
         }]*/
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
            img += '<span class="rate">' + (item[item.res] * 100).toFixed(2) + '%</span>'
            img += '</label>'
            img += '<label class="datetime ">'
            img += '<span class="time">' + item.created_at + '</span>'
            img += '</label>'
            img += '</div>'
            img += '<div class="imgbox">'
            img += '<img class="imgview" alt="..." src=' + item.url + '   />'
            img += '</div>'
            img += '</div>'
            div.append(img)
        })
    }

    function pageselectCallback(page_index, jq) {
        if (!isScreen) {
            if (isFirstLoad || page_index != 0) {
                $.ajax({
                    type: 'get',
                    url: defaults.viewUrl,
                    data: {
                        "username": username,
                        "time": time,
                        "type": type,
                        "sort": sort,
                        "sord": sord,
                        "review": review,
                        "isAll": isAll,
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
                        initImg(data)
                    },
                    complete: function () {
                        $.LoadingOverlay("hide");
                    },
                    error: function (data) {
                    },
                });
                isFirstLoad = true
            }
        } else {
            if (isFirstLoad || page_index != 0) {
                $.ajax({
                    type: 'get',
                    url: defaults.screenUrl,
                    data: {
                        "username": username,
                        "time": time,
                        "sort": sort,
                        "sord": sord,
                        "screen": "isScreen",
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
                        initImg(data)
                    },
                    complete: function () {
                        $.LoadingOverlay("hide");
                    },
                    error: function (data) {
                    },
                });
                isFirstLoad = true
            }
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
    function initPhoto() {
        isFirstLoad = false
        $.ajax({
            type: 'get',
            url: defaults.viewUrl,
            data: {
                "username": username,
                "time": time,
                "type": type,
                "review": review,
                "sort": sort,
                "sord": sord,
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
                $("#type" + type + _review).addClass("button-select");
                $("#datetimepicker").attr("value", time)
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

    $('#autocomplete-search').autocomplete({
        serviceUrl: '/stat_search_users',
        params: {
            page: "1",
            rows: "10",
        },
        ajaxSettings: {
            headers: {
                "LoginToken": loginToken,
                "Access-Control-Allow-Headers": "LoginToken",
            },
        },
        onSelect: function (suggestion) {
            username = suggestion.data
            initPhoto()
        },
        transformResult: function (response) {
            response = $.parseJSON(response);
            var result = []
            response.result.forEach(function (item, idx) {
                var resultTemp = {}
                resultTemp = {
                    "username": item.username,
                    "company": item.company
                }
                result.push(resultTemp)
                if (item.company == "") {
                    resultTemp = {
                        "username": item.username,
                        "company": item.username
                    }
                    result.push(resultTemp)
                }
            })
            return {
                suggestions: $.map(result, function (dataItem) {
                    return {value: dataItem.company, data: dataItem.username};
                })
            }
        }
    });

    $("#search").click(function(data){
        var searchtext = $("#autocomplete-search").val()
        $.ajax({
            type: 'get',
            url: defaults.searchUrl,
            data:{
                query: searchtext,
                page:"1",
                rows:"10",

            },
            headers: {
                "LoginToken": loginToken,
                "Access-Control-Allow-Headers": "LoginToken",
            },
            success: function(data) {
                window.location.href="/searchuser?url=lookimg&text="+$("#autocomplete-search").val()
            },
            error: function (data) {
                //window.location.href="/searchuser "
            },
        });
    });

    $("#up").click(function () {
        isFirstLoad = false
        sord = "desc"
        sort = ""
        if (isScreen) {
            $.ajax({
                type: 'get',
                url: defaults.screenUrl,
                data: {
                    "username": username,
                    "time": time,
                    "screen": "isScreen",
                    "sort": sort,
                    "sord": sord,
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
                    $("#up").hide();
                    $("#down").show();
                    initPagination(data)
                    initImg(data)
                },
                complete: function () {
                    $.LoadingOverlay("hide");
                },
                error: function (data) {
                },
            });
        } else {
            $.ajax({
                type: 'get',
                url: defaults.viewUrl,
                data: {
                    "username": username,
                    "time": time,
                    "type": type,
                    "review": review,
                    "isAll": isAll,
                    "sort": sort,
                    "sord": sord,
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
                    $("#up").hide();
                    $("#down").show();
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
    });

    $("#down").click(function () {
        isFirstLoad = false
        sord = "asc"
        sort = ""
        if (isScreen) {
            $.ajax({
                type: 'get',
                url: defaults.screenUrl,
                data: {
                    "username": username,
                    "time": time,
                    "screen": "isScreen",
                    "sort": sort,
                    "sord": sord,
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
                    $("#down").hide();
                    $("#up").show();
                    initPagination(data)
                    initImg(data)
                },
                complete: function () {
                    $.LoadingOverlay("hide");
                },
                error: function (data) {
                },
            });
        } else {
            $.ajax({
                type: 'get',
                url: defaults.viewUrl,
                data: {
                    "username": username,
                    "time": time,
                    "type": type,
                    "review": review,
                    "isAll": isAll,
                    "sort": sort,
                    "sord": sord,
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
                    $("#down").hide();
                    $("#up").show();
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
    });

    $("#timeup").click(function () {
        isFirstLoad = false
        sord = "desc"
        sort = "created_at"
        if (isScreen) {
            $.ajax({
                type: 'get',
                url: defaults.screenUrl,
                data: {
                    "username": username,
                    "time": time,
                    "screen": "isScreen",
                    "sort": sort,
                    "sord": sord,
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
        } else {
            $.ajax({
                type: 'get',
                url: defaults.viewUrl,
                data: {
                    "username": username,
                    "time": time,
                    "type": type,
                    "review": review,
                    "isAll": isAll,
                    "sort": sort,
                    "sord": sord,
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
        }
    });

    $("#timedown").click(function () {
        isFirstLoad = false
        sord = "asc"
        sort = "created_at"
        if (isScreen) {
            $.ajax({
                type: 'get',
                url: defaults.screenUrl,
                data: {
                    "username": username,
                    "time": time,
                    "screen": "isScreen",
                    "sort": sort,
                    "sord": sord,
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
                    initImg(data)
                },
                complete: function () {
                    $.LoadingOverlay("hide");
                },
                error: function (data) {
                },
            });
        } else {
            $.ajax({
                type: 'get',
                url: defaults.viewUrl,
                data: {
                    "username": username,
                    "time": time,
                    "type": type,
                    "review": review,
                    "isAll": isAll,
                    "sort": sort,
                    "sord": sord,
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
                    initImg(data)
                },
                complete: function () {
                    $.LoadingOverlay("hide");
                },
                error: function (data) {
                },
            });
        }
    });

    $("#types").click(function () {
        $("#types").addClass("button-select");
        $("#typep").removeClass("button-select");
        $("#type_screen").removeClass("button-select");
        $(".reviewcheck").show();
        $(".showreview").prop('checked', false);
        $(".checktype").prop('checked', true);
        initCheckBox(".showreview")
        initCheckBox(".checktype")
        isFirstLoad = false
        type = "s"
        review = false
        isScreen = false
        isAll = false
        pBoth = 1
        sBoth = 1
        $.ajax({
            type: 'get',
            url: defaults.viewUrl,
            data: {
                "username": username,
                "time": time,
                "type": type,
                "review": review,
                "sort": sort,
                "sord": sord,
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
                initImg(data)
            },
            complete: function () {
                $.LoadingOverlay("hide");
            },
            error: function (data) {
            },

        });
    });

    $("#typep").click(function () {
        $("#types").removeClass("button-select");
        $("#typep").addClass("button-select");
        $("#type_screen").removeClass("button-select");
        $(".reviewcheck").show();
        $(".showreview").prop('checked', false);
        $(".checktype").prop('checked', true);
        initCheckBox(".showreview")
        initCheckBox(".checktype")
        isFirstLoad = false
        type = "p"
        pBoth = 1
        sBoth = 1
        review = false
        isAll = false
        isScreen = false
        $.ajax({
            type: 'get',
            url: defaults.viewUrl,
            data: {
                "username": username,
                "time": time,
                "type": type,
                "review": review,
                "sort": sort,
                "sord": sord,
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
                initImg(data)
            },
            complete: function () {
                $.LoadingOverlay("hide");
            },
            error: function (data) {
            },

        });
    });

    $("#type_screen").click(function () {
        $("#types").removeClass("button-select");
        $("#typep").removeClass("button-select");
        $("#type_screen").addClass("button-select");
        $(".reviewcheck").hide();
        isScreen = true
        isFirstLoad = false
        $.ajax({
            type: 'get',
            url: defaults.screenUrl,
            data: {
                "username": username,
                "time": time,
                "screen": "isScreen",
                "sort": sort,
                "sord": sord,
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
                initImg(data)
            },
            complete: function () {
                $.LoadingOverlay("hide");
            },
            error: function (data) {
            },

        });
    });
    initCheckBox(".showreview")
    initCheckBox(".checktype")
    function initCheckBox(dom) {
        if ($(dom).is(':checked')) {
            $(dom).prop("disabled", true);
        } else {
            $(dom).prop("disabled", false);
        }
    }

    function changeCheckBoxStatus(dom, able) {
        $(dom).prop("disabled", able);
    }

    $(".showreview").change(function () {
        isFirstLoad = false
        if ($(".showreview").is(':checked')) {
            $(".checktype").attr("checked", false);
            $(".showreview").prop("disabled", true);
            $(".checktype").prop("disabled", false);
            review = true
        }
        $.ajax({
            type: 'get',
            url: defaults.viewUrl,
            data: {
                "username": username,
                "time": time,
                "type": type,
                "review": review,
                "isAll": isAll,
                "sort": sort,
                "sord": sord,
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
                initImg(data)
            },
            complete: function () {
                $.LoadingOverlay("hide");
            },
            error: function (data) {
            },

        });
    })

    $(".checktype").change(function () {
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
                "username": username,
                "time": time,
                "type": type,
                "review": review,
                "isAll": isAll,
                "sort": sort,
                "sord": sord,
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
                initImg(data)
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
            isFirstLoad = false
            time = $input.val();
            $.ajax({
                type: 'get',
                url: defaults.viewUrl,
                data: {
                    "username": username,
                    "time": time,
                    "sort": sort,
                    "sord": sord,
                    "review": review,
                    "page": defaults.page,
                    "rows": defaults.rows,
                    "type": type,
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
                    initImg(data)
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
