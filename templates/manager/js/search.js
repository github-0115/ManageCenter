function getDetail(id){
    var rowData = $("#jqGrid").jqGrid('getRowData',id);
    function getUrlParam(name) {
        var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
        var r = window.location.search.substr(1).match(reg);  //匹配目标参数
        if (r != null) return unescape(r[2]);
        return null; //返回参数值
    }
    url = getUrlParam("url");
    console.info(url)
    console.info("=====")
    console.info("/"+url+"?username="+rowData.username)
   window.location.href="/"+url+"?username="+rowData.username
}
$(document).ready(function() {
    $(function() {
        pageInit();
    });

    function pageInit() {
        $("#jqGrid").jqGrid({
            url: "/stat_search_users",
            mtype: "GET",
            colModel: [{
                    label: '用户名',
                    name: 'username',
                    align: 'center',
                },
                {
                    label: '使用天数',
                    name: 'daynum',
                    align: 'center',
                },
                {
                    label: '公司',
                    name: 'company',
                    align: 'center',
                },
                {
                    label: '联系电话',
                    name: 'phonenum',
                    align: 'center',

                }, {
                    label: '联系邮箱',
                    name: 'email',
                    align: 'center'

                },  {
                    name: '查看详情',
                    align: 'center',  
                    //formatter:'actions',
                    formatter: function (value, grid, rows, state) {
                        return "<a style=\"color:#f60;cursor:pointer\" onclick=\"getDetail(" + grid.rowId + ")\">查看详情</a>"
                    }
                }
            ],
            prmNames: {
                page: "page", // 表示请求页码的参数名称
                rows: "rows", // 表示请求行数的参数名称
                sort: "sort",
                search: "search",
            },
             postData:{query:getUrlParam("text")},

            ajaxGridOptions: {
                headers: {
                    "LoginToken": loginToken,
                    "Access-Control-Allow-Headers": "LoginToken",
                },
            },
            jsonReader: {
                root: "result",
                page: "page",
                total: "total",
                records: "record",
                repeatitems: false
            },
            rowList: [15, 20, 25],
            viewrecords: true, 
            autowidth:true,
       
            height: 600,
            rowNum: 15,
            datatype: "json",
            pager: "#jqGridPager",
            autosearch: true,
           
            altRows: true,

            onCellSelect:function(rowid,iCol){
                var rowData = getRowData(rowid);
                var username =rowData.username;
            },
            beforeProcessing: function (data) {
            },
            loadComplete : function() {
                var table = this;
                setTimeout(function(){
                          updatePagerIcons(table);
                }, 0);
            },
        });

      function updatePagerIcons(table) {
                    var replacement = 
                    {
                        'ui-icon-seek-first' : 'icon-double-angle-left bigger-140',
                        'ui-icon-seek-prev' : 'icon-angle-left bigger-140',
                        'ui-icon-seek-next' : 'icon-angle-right bigger-140',
                        'ui-icon-seek-end' : 'icon-double-angle-right bigger-140'
                    };
                    $('.ui-pg-table:not(.navtable) > tbody > tr > .ui-pg-button > .ui-icon').each(function(){
                        var icon = $(this);
                        var $class = $.trim(icon.attr('class').replace('ui-icon', ''));
                        
                        if($class in replacement) icon.attr('class', 'ui-icon '+replacement[$class]);
                    })
                }

        function formatTitle(cellValue, options, rowObject) {
            return cellValue.substring(0, 50) + "...";
        };

        function formatLink(cellValue, options, rowObject) {
            return "<a href='" + cellValue + "'>" + cellValue.substring(0, 25) + "..." + "</a>";
        };

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
       function getUrlParam(name) {
        var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
        var r = window.location.search.substr(1).match(reg);  //匹配目标参数
        if (r != null) return unescape(r[2]);
        return null; //返回参数值
    }
});
