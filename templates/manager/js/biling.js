var currbank = "中国工商银行"
var currdiscount = "8.5"

$("#bank").change(function () {
    var value = $("#bank").val();
    currbank = value;
});

function biling(id) {
    var defaults = {
        bilingUrl: "/bill/one_bills",
    };
    var rowData = $("#jqGrid").jqGrid('getRowData', id);
    var html = ""
    html += '<div class="apisetbox"> ' +
        '<div class="apisetting"><p class="title">用户名:</p><span id="username">' + rowData.username + '</span></div> ' +
        '<div class="apisetting"><p class="title">使用总量:</p><span id="total"></span></div> ' +
      /*  '<div class="apisetting"><p class="title">优惠区间：</p><select id="discount" > ' +
        '<option value="1">无优惠</option>' +
        '<option value="0.9">9折</option>' +
        '<option value="0.85">8.5折</option> ' +
        '<option value="0.75">7.5折</option>' +
        '<option value="0.7">7折</option>' +
        '<option value="0.65">6.5折</option>' +
        '</select></div> ' +*/
        '<div class="apisetting" ><p class="title" id="type">结算类型：</p> <span class="type"></span></div> ' +
        '<div class="apisetting" ><p class="title" id="circle">结算周期：</p><span class="period"></span> </div> ' +
        '<div class="apisetting" id="billdate" style="display: none"><p style="color: red">（暂时还未到结算时间）</p> </div> ' +
        '<div class="apisetting"><p class="title">原收费用：</p><span class="p-num" id="orisum"></span> <span> 元</span></div> ' +
        '<div class="apisetting"><p class="title">实收费用：</p><input class="p-num" id="nowsum"><span> 元</span></div> ' +
        '</div>' +
        '<style>.layui-layer-btn .layui-layer-btn0 {border-color: #4898d5;background-color: #2e8ded;color: #fff;display: block;margin: auto;width: 125px;text-align: center;}</style>'

    GetApisetting()

    function GetApisetting() {
        $.ajax({
            type: 'get',
            url: defaults.bilingUrl,
            data: {
                "username": rowData.username,
                "finalbill_id": rowData.finalbill_id,
            },
            headers: {
                "LoginToken": loginToken,
                "Access-Control-Allow-Headers": "LoginToken",
            },
            beforeSend: function () {

            },
            success: function (data) {
                $("#total").text(rowData.used_total + '次'),
                $(".type").text(rowData.bill_type),
                $(".period").text(rowData.period),
                $("#orisum").text(rowData.total_fee),
                $("#nowsum").val(rowData.total_fee)
                console.info(data.bill.billdate)
                if(data.billdate==0){
                    $('#billdate').show()
                }else {
                    $('#billdate').hide()
                }
            },
            complete: function () {

            },
            error: function (data) {
            },
        });
    }

    layer.open({
        type: 1,
        skin: 0,
        area: ['370px', '360px'], //宽高
        shade: 0.2,
        title: [
            '结算处理',
            'background-color:#0b6cbc; color:#fff;'
        ],
        content: html
        , btn: ['生成结算订单']
        , yes: function (index, layero) {
            $.ajax({
                type: 'put',
                url: "/bill/favor_bills",
                data: JSON.stringify({
                    "username": rowData.username,
                    finalbill_id: rowData.finalbill_id,
                   /* favor_range: $("#discount").val(),*/
                    paid_amount: parseInt($("#nowsum").val(), 10),
                }),
                headers: {
                    "LoginToken": loginToken,
                    "Access-Control-Allow-Headers": "LoginToken",
                },
                success: function (data) {
                    layer.msg('保存成功', {
                        icon: 1
                    });
                    $("#total").text(rowData.used_total + '次'),
                    $(".type").text(rowData.bill_type),
                    $(".period").text(rowData.period),
                    $("#orisum").val(rowData.total_fee)
                    $("#nowsum").val(rowData.paid_amount)

                    GetApisetting()
                    $("#jqGrid").jqGrid().trigger("reloadGrid");
                    layer.close(index);
                },
                error: function (data) {
                    layer.msg('保存失败', {
                        icon: 5
                    });
                },
            });

        },
        cancel: function (index) {
            layer.close(index);
        }
    });


}
function verify(id) {
    var rowData = $("#jqGrid").jqGrid('getRowData', id);
    var html = ""
    html += '<div class="apisetbox"> ' +
        '<div class="apisetting"><p class="title">用户名:</p><span id="username">' + rowData.username + '</span></div> ' +
        '<div class="apisetting"><p class="title">使用总量:</p><span id="total">' + rowData.total + '</span></div> ' +
        '<div class="apisetting"><p class="title">实收费用：</p><span id="nowsum"> 元</span></div> ' +
       /* '<div class="apisetting"><p class="title">转账银行：</p><span id="bank"></span></div> ' +*/
        '<div class="apisetting"><p class="title">转账银行：</p><select id="bank"> ' +
        '<option value="中国工商银行">中国工商银行</option> ' +
        '<option value="中国建设银行">中国建设银行</option>' +
        '<option value="中国银行">中国银行</option>' +
        '<option value="中国农业银行">中国农业银行</option>' +
        '<option value="招商银行">招商银行</option>' +
        '</select></div> ' +
        '<div class="apisetting"><p class="title">转账账号：</p><input id="account" ></div> ' +
        '<div class="apisetting"><p class="title">转账单号：</p><input id="num"></div> ' +
        '</div>'

    Getbanksetting()

    function Getbanksetting() {
        $.ajax({
            type: 'get',
            url: "/bill/one_bills",
            data: {
                "username": rowData.username,
                "finalbill_id": rowData.finalbill_id,
            },
            headers: {
                "LoginToken": loginToken,
                "Access-Control-Allow-Headers": "LoginToken",
            },
            beforeSend: function () {

            },
            success: function (data) {
                $("#total").text(data.bill.used_total + '次'),
                $("#nowsum").text(rowData.paid_amount+'元')
                $("#bank").val(rowData.transfer_bank)
                $("#account").val(data.bill.transfer_account)
                $("#num").val(data.bill.transfer_id)
            },
            complete: function () {

            },
            error: function (data) {
            },
        });
    }

    layer.open({
        type: 1,
        skin: 0,
        area: ['370px', '370px'], //宽高
        shade: 0.2,
        title: [
            '完善信息',
            'background-color:#0b6cbc; color:#fff;'
        ],
        content: html
        , btn: ['确认收到款','没有收到这笔款项']
        , yes: function (index, layero) {
            $.ajax({
                type: 'put',
                url: "/bill/transfers_bills",
                data: JSON.stringify({
                    "username": rowData.username,
                    "finalbill_id": rowData.finalbill_id,
                    "transfer_id": $("#num").val(),
                    transfer_bank: $("#bank").val(),
                    transfer_account: $("#account").val()
                }),
                headers: {
                    "LoginToken": loginToken,
                    "Access-Control-Allow-Headers": "LoginToken",
                },
                success: function (data) {

                    $("#total").text(rowData.used_total + '次'),
                    $("#nowsum").text(rowData.paid_amount)
                    $("#num").val(rowData.transfer_id)
                    $("#bank").val(rowData.transfer_bank)
                    $("#account").val(rowData.transfer_account)

                        layer.msg('确认收款成功', {
                            icon: 1
                        });

                    Getbanksetting();
                    $("#jqGrid").jqGrid().trigger("reloadGrid");
                    layer.close(index);
                },
                error: function (data) {
                    layer.msg('确认收款失败', {
                        icon: 5
                    });
                    var status = {
                        709:"银行账号格式不对，请重新输入",
                        710:"这是一个无效的银行账号",
                    };
                    layer.msg(status[data.responseJSON.code] ,{
                        icon: 5,
                    });
                },
            });

        },
        btn2:function(index, layero){
            $.ajax({
                type: 'post',
                url: "/bill/not_pass_bills",
                data: JSON.stringify({
                    "username": rowData.username,
                }),
                headers: {
                    "LoginToken": loginToken,
                    "Access-Control-Allow-Headers": "LoginToken",
                },
                success: function (data) {
                    layer.msg('已成功通知用户未收到款项信息', {
                        icon: 1
                    });

                    Getbanksetting();
                    $("#jqGrid").jqGrid().trigger("reloadGrid");
                    layer.close(index);
                },
                error: function (data) {
                    layer.msg('通知失败', {
                        icon: 5
                    });
                },
            });

        },
        cancel: function (index) {
            layer.close(index);
        }
    });
}
function detail(id) {
    var rowData = $("#jqGrid").jqGrid('getRowData', id);
    window.localStorage.setItem("finalbill_id", rowData.finalbill_id);
    window.localStorage.setItem("username", rowData.username);
    layer.open({
        type: 2,
        skin: 0,
        area: ['600px', '400px'], //宽高
        shade: 0.2,
        title: [
            '查看详情',
            'background-color:#0b6cbc; color:#fff;'
        ],
        content: '/billdetail',
        end:function(){
            window.localStorage.removeItem("username");
            window.localStorage.removeItem("finalbill_id");
        }
    });
}

$(document).ready(function () {
    $(function () {
        pageInit();
    });

    function pageInit() {
        $("#jqGrid").jqGrid({
            url: "/bill/all_bills",
            mtype: "GET",
            colModel: [{
                label: '用户名',
                name: 'username',
                align: 'center',
                width: '90'
            }, {
                label: 'finalbill_id',
                name: 'finalbill_id',
                align: 'center',
                hidden:true
            },
                {
                    label: '使用总量',
                    name: 'used_total',
                    align: 'center',
                    width: '90'
                }, {
                    label: '消费总额',
                    name: 'total_fee',
                    align: 'center',
                    width: '90'
                },{
                    label: '实收价格',
                    name: 'paid_amount',
                    align: 'center',
                    width: '90'
                }, {
                    label: '优惠区间',
                    name: 'favor_range',
                    align: 'center',
                    width: '60',
                    hidden: true
                }, {
                    label: '结算周期',
                    name: 'period',
                    align: 'center',
                }, {
                    label: '转账单号',
                    name: 'transfer_id',
                    align: 'center',
                    hidden: true
                }, {
                    label: '银行名称',
                    name: 'transfer_bank',
                    align: 'center',
                    hidden: true
                }, {
                    label: '银行账号',
                    name: 'transfer_account',
                    align: 'center',
                    hidden: true
                }, {
                    label: '实付金额',
                    name: 'paid_amount',
                    align: 'center',
                    hidden: true
                }, {
                    label: '状态码',
                    name: 'status',
                    align: 'center',
                    hidden: true
                }, {
                    label: '类型',
                    name: 'bill_type',
                    align: 'center',
                    width: '60'
                }, {
                    name: '操作', index: '',
                    align: 'center',
                    width: '90',
                    formatter: function (value, grid, rows, state) {
                        if(rows.bill_type=="月结")
                        return "<a style=\"color:#f60;cursor:pointer\" onclick=\"biling(" + grid.rowId + ")\">处理</a>"
                        else{
                            return "<a style=\"color:#f60;cursor:pointer\" onclick=\"biling(" + grid.rowId + ")\">处理</a>&nbsp;" +
                            "<a style=\"color:#f60;cursor:pointer\" onclick=\"detail(" + grid.rowId + ")\">详情</a>"
                        }
                    }
                },
            ],
            prmNames: {
                page: "page", // 表示请求页码的参数名称
                rows: "rows", // 表示请求行数的参数名称
                sort: "sort",
                search: "search",
            },
            postData:{
                status:"0",
            },
            ajaxGridOptions: {
                headers: {
                    "LoginToken": loginToken,
                    "Access-Control-Allow-Headers": "LoginToken",
                },
            },
            jsonReader: {
                root: "bill",
                page: "page",
                total: "total",
                records: "record",
                repeatitems: false
            },
            rowList: [15, 20, 25],
            viewrecords: true,
            autowidth: true,

            height: 600,
            rowNum: 15,
            datatype: "json",
            pager: "#jqGridPager",
            autosearch: true,

            altRows: true,
            beforeProcessing: function (data) {
                var bills=[]
                data.bill.forEach(function(item){
                    if(item.status=="0"){
                        var start=new Date(item.period_start)
                        var end=new Date(item.period_end)
                        item["period"]=start.format("yyyy-MM-dd")+' 至 '+end.format("yyyy-MM-dd")
                        bills.push(item)
                    }
                })
                data.bill=bills
            },
            loadComplete: function () {
                var table = this;
                setTimeout(function () {
                    updatePagerIcons(table);
                }, 0);
            },
        });
        $("#jqGrid").jqGrid('navGrid', '#jqGridPager', {
            edit: false,
            add: false,
            del: false,
            search: true,
            searchicon: 'icon-search orange',
            refresh: true,
            refreshicon: 'icon-refresh green',
            view: true,
            viewicon: 'icon-zoom-in grey',

        });
        /* jQuery("#jqGrid").navButtonAdd('#jqGridPager',{
         caption:"Excel",
         buttonicon:"ui-icon-excel",
         onClickButton: function(){
         alert("导出excel");
         },
         position:"last"
         });*/
    }
});

function updatePagerIcons(table) {
    var replacement =
    {
        'ui-icon-seek-first': 'icon-double-angle-left bigger-140',
        'ui-icon-seek-prev': 'icon-angle-left bigger-140',
        'ui-icon-seek-next': 'icon-angle-right bigger-140',
        'ui-icon-seek-end': 'icon-double-angle-right bigger-140'
    };
    $('.ui-pg-table:not(.navtable) > tbody > tr > .ui-pg-button > .ui-icon').each(function () {
        var icon = $(this);
        var $class = $.trim(icon.attr('class').replace('ui-icon', ''));

        if ($class in replacement) icon.attr('class', 'ui-icon ' + replacement[$class]);
    })
}

function formatTitle(cellValue, options, rowObject) {
    return cellValue.substring(0, 50) + "...";
};

function formatLink(cellValue, options, rowObject) {
    return "<a href='" + cellValue + "'>" + cellValue.substring(0, 25) + "..." + "</a>";
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

$("#beforedeal").click(function () {
    $("#jqGrid").jqGrid('GridUnload');
    $("#jqGrid").jqGrid({
        url: "/bill/all_bills",
        mtype: "GET",
        colModel: [{
            label: '用户名',
            name: 'username',
            align: 'center',
            width: '90'
        }, {
            label: 'bill_id',
            name: 'finalbill_id',
            align: 'center',
            hidden:true
        },
            {
                label: '使用总量',
                name: 'used_total',
                align: 'center',
                width: '90'
            }, {
                label: '消费总额',
                name: 'total_fee',
                align: 'center',
                width: '90'
            }, {
                label: '实收价格',
                name: 'paid_amount',
                align: 'center',
                width: '90'
            }, {
                label: '优惠区间',
                name: 'favor_range',
                align: 'center',
                width:'60',
                hidden: true
            }, {
                label: '结算周期',
                name: 'period',
                align: 'center',
            },
            {
                label: '转账单号',
                name: 'transfer_id',
                align: 'center',
                hidden: true
            }, {
                label: '银行名称',
                name: 'transfer_bank',
                align: 'center',
                hidden: true
            }, {
                label: '银行账号',
                name: 'transfer_account',
                align: 'center',
                hidden: true
            }, {
                label: '实付金额',
                name: 'paid_amount',
                align: 'center',
                hidden: true
            }, {
                label: '类型',
                name: 'bill_type',
                align: 'center',
                width:'60'
            },{
                label: '状态码',
                name: 'status',
                align: 'center',
                hidden: true
            },{
                name: '操作',
                index: '',
                align: 'center',
                width: '90',
                formatter: function (value, grid, rows, state) {
                    if(rows.bill_type=="月结")
                        return "<a style=\"color:#f60;cursor:pointer\" onclick=\"biling(" + grid.rowId + ")\">处理</a>"
                    else{
                        return "<a style=\"color:#f60;cursor:pointer\" onclick=\"biling(" + grid.rowId + ")\">处理</a>&nbsp;" +
                            "<a style=\"color:#f60;cursor:pointer\" onclick=\"detail(" + grid.rowId + ")\">详情</a>"
                    }
                }
            },
        ],
        prmNames: {
            page: "page", // 表示请求页码的参数名称
            rows: "rows", // 表示请求行数的参数名称
            sort: "sort",

            search: "search",
        },
        postData:{
            status:"0",
        },
        ajaxGridOptions: {
            headers: {
                "LoginToken": loginToken,
                "Access-Control-Allow-Headers": "LoginToken",
            },
        },
        jsonReader: {
            root: "bill",
            page: "page",
            total: "total",
            records: "record",
            repeatitems: false
        },
        rowList: [15, 20, 25],
        viewrecords: true,
        autowidth: true,

        height: 600,
        rowNum: 15,
        datatype: "json",
        pager: "#jqGridPager",
        autosearch: true,

        altRows: true,
        beforeProcessing: function (data) {
            var bills=[]
            data.bill.forEach(function(item){
                if(item.status=="0"){
                    var start=new Date(item.period_start)
                    var end=new Date(item.period_end)
                    item["period"]=start.format("yyyy-MM-dd")+' 至 '+end.format("yyyy-MM-dd")
                    bills.push(item)
                }
            })
            data.bill=bills
        },
        loadComplete: function () {
            var table = this;
            setTimeout(function () {
                updatePagerIcons(table);
            }, 0);
        },
    });
    $("#jqGrid").jqGrid('navGrid', '#jqGridPager', {
        edit: false,
        add: false,
        del: false,
        search: true,
        searchicon: 'icon-search orange',
        refresh: true,
        refreshicon: 'icon-refresh green',
        view: true,
        viewicon: 'icon-zoom-in grey',

    });
})

$("#afterdeal").click(function () {
    $("#jqGrid").jqGrid('GridUnload');
    $("#jqGrid").jqGrid({
        url: "/bill/all_bills",
        mtype: "GET",
        colModel: [{
            label: '用户名',
            name: 'username',
            align: 'center',
            width: '90'
        },
            {
                label: '使用总量',
                name: 'used_total',
                align: 'center',
                width: '90'
            }, {
                label: '消费总额',
                name: 'total_fee',
                align: 'center',
                width: '90'
            },{
                label: '实收价格',
                name: 'paid_amount',
                align: 'center',
                width: '90'
            },  {
                label: '订单号',
                name: 'finalbill_id',
                align: 'center',
            },{
                label: '优惠区间',
                name: 'favor_range',
                align: 'center',
                width:'60',
                hidden:true
            }, {
                label: '结算周期',
                name: 'period',
                align: 'center',
            }, {
                label: '转账单号',
                name: 'transfer_id',
                align: 'center',
                hidden: true
            }, {
                label: '银行名称',
                name: 'transfer_bank',
                align: 'center',
                hidden: true
            }, {
                label: '银行账号',
                name: 'transfer_account',
                align: 'center',
                hidden: true
            }, {
                label: '实付金额',
                name: 'paid_amount',
                align: 'center',
                hidden: true
            },{
                label: '状态码',
                name: 'status',
                align: 'center',
                hidden: true
            },
            {
                label: '状态',
                name: 'remark',
                align: 'center',
                hidden:true
            },
            {
                label: '类型',
                name: 'bill_type',
                align: 'center',
                width:'60'
            }, {
                name: '操作',
                index: '',
                align: 'center',
                width: '90',
                formatter: function (value, grid, rows, state) {
                    if(rows.transfer_bank=="")
                        return "待支付&nbsp;<a class=\" btn-info btn btn-xs\"  style=\"color:#f60;cursor:pointer\" onclick=\"verify(" + grid.rowId + ")\">银行转账</a>"
                    else{
                        return "已支付&nbsp;<a class=\" btn-info btn btn-xs\"  style=\"color:#f60;cursor:pointer\" onclick=\"verify(" + grid.rowId + ")\">银行转账</a>"
                    }
                }
            },
        ],
        prmNames: {
            page: "page", // 表示请求页码的参数名称
            rows: "rows", // 表示请求行数的参数名称
            sort: "sort",
            search: "search",
        },
        postData:{
            status:"1",

        },

        ajaxGridOptions: {
            headers: {
                "LoginToken": loginToken,
                "Access-Control-Allow-Headers": "LoginToken",
            },
        },
        jsonReader: {
            root: "bill",
            page: "page",
            total: "total",
            records: "record",
            repeatitems: false
        },
        rowList: [15, 20, 25],
        viewrecords: true,
        autowidth: true,

        height: 600,
        rowNum: 15,
        datatype: "json",
        pager: "#jqGridPager",
        autosearch: true,

        altRows: true,
        beforeProcessing: function (data) {
            var bills=[]
            data.bill.forEach(function(item){
                if(item.status=="1"){
                    var start=new Date(item.period_start)
                    var end=new Date(item.period_end)
                    item["period"]=start.format("yyyy-MM-dd")+' 至 '+end.format("yyyy-MM-dd")
                    bills.push(item)
                }
            })
            data.bill=bills
        },
        loadComplete: function () {
            var table = this;
            setTimeout(function () {
                updatePagerIcons(table);
            }, 0);
        },
    });
    $("#jqGrid").jqGrid('navGrid', '#jqGridPager', {
        edit: false,
        add: false,
        del: false,
        search: true,
        searchicon: 'icon-search orange',
        refresh: true,
        refreshicon: 'icon-refresh green',
        view: true,
        viewicon: 'icon-zoom-in grey',

    });
})

$("#over").click(function () {
    $("#jqGrid").jqGrid('GridUnload');
    $("#jqGrid").jqGrid({
        url: "/bill/all_bills",
        mtype: "GET",
        colModel: [{
            label: '用户名',
            name: 'username',
            align: 'center',
            width: '90'
        },
            {
                label: '使用总量',
                name: 'used_total',
                align: 'center',
                width: '90'
            }, {
                label: '消费总额',
                name: 'total_fee',
                align: 'center',
                width: '90'
            }, {
                label: '实收价格',
                name: 'paid_amount',
                align: 'center',
                width: '90'
            },{
                label: '订单号',
                name: 'finalbill_id',
                align: 'center',
            }, {
                label: '优惠区间',
                name: 'favor_range',
                align: 'center',
                width: '60',
                hidden:true
            }, {
                label: '结算周期',
                name: 'period',
                align: 'center',
            }, {
                label: '转账单号',
                name: 'transfer_id',
                align: 'center',
                hidden: true
            }, {
                label: '银行名称',
                name: 'transfer_bank',
                align: 'center',
                hidden: true
            }, {
                label: '银行账号',
                name: 'transfer_account',
                align: 'center',
                hidden: true
            }, {
                label: '实付金额',
                name: 'paid_amount',
                align: 'center',
                hidden: true
            },
            {
                label: '类型',
                name: 'bill_type',
                align: 'center',
                width:'60'
            },{
                label: '状态码',
                name: 'status',
                align: 'center',
                hidden: true
            },
            {
                label: '状态',
                name: 'remark',
                align: 'center',
                width:'250'
            }, {
                label: '交易时间',
                name: 'transaction_time',
                align: 'center',
                width: '90'
            },
        ],
        prmNames: {
            page: "page", // 表示请求页码的参数名称
            rows: "rows", // 表示请求行数的参数名称
            sort: "sort",
            search: "search",
        },
        postData:{
            status:"2",
        },
        ajaxGridOptions: {
            headers: {
                "LoginToken": loginToken,
                "Access-Control-Allow-Headers": "LoginToken",
            },
        },
        jsonReader: {
            root: "bill",
            page: "page",
            total: "total",
            records: "record",
            repeatitems: false
        },
        rowList: [15, 20, 25],
        viewrecords: true,
        autowidth: true,

        height: 600,
        rowNum: 15,
        datatype: "json",
        pager: "#jqGridPager",
        autosearch: true,

        altRows: true,
        beforeProcessing: function (data) {
            var bills=[]
            data.bill.forEach(function(item){
                if(item.status=="2"){
                    var start=new Date(item.period_start)
                    var end=new Date(item.period_end)
                    item["period"]=start.format("yyyy-MM-dd")+' 至 '+end.format("yyyy-MM-dd")
                    var transaction_time=new Date(item.transaction_time)
                    item["transaction_time"]=transaction_time.format("yyyy-MM-dd")
                    item["remark"]='已通过'+ item.transfer_bank+'转账,凭证号'+item.transfer_id
                    bills.push(item)

                }

            })
            data.bill=bills
        },
        loadComplete: function () {
            var table = this;
            setTimeout(function () {
                updatePagerIcons(table);
            }, 0);
        },
    });
    $("#jqGrid").jqGrid('navGrid', '#jqGridPager', {
        edit: false,
        add: false,
        del: false,
        search: true,
        searchicon: 'icon-search orange',
        refresh: true,
        refreshicon: 'icon-refresh green',
        view: true,
        viewicon: 'icon-zoom-in grey',

    });
})



