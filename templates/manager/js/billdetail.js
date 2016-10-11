/**
 * Created by huang on 2016/9/20.
 */

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

$(document).ready(function () {
    $(function () {
        pageInit();
    });

    finalbill_id = window.localStorage.getItem("finalbill_id")
    username = window.localStorage.getItem("username")

    function pageInit() {
        $("#jqGrid").jqGrid({
            url: "/bill/detail_bills",
            mtype: "GET",
            colModel: [{
                label: '用户名',
                name: 'username',
                align: 'center',
                width: '90',
                hidden:true
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
                    width: '90',
                    hidden:true
                }, {
                    label: '优惠区间',
                    name: 'favor_range',
                    align: 'center',
                    width: '60',
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
                    width: '60',
                    hidden:true
                },
            ],
            prmNames: {
                page: "page", // 表示请求页码的参数名称
                rows: "rows", // 表示请求行数的参数名称
                sort: "sort",
                search: "search",
            },
            postData:{
                "username": username,
                "finalbill_id": finalbill_id,
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

            rowList: [6, 12, 24],
            viewrecords: true,
            autowidth: true,

            height: 220,
            width:500,
            rowNum: 12,
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

