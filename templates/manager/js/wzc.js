$(document).ready(function() {
    $(function() {
        pageInit();
    });

    function pageInit() {
        $("#jqGrid").jqGrid({
            url: "/potential_users",
            mtype: "GET",
            colModel: [
                {
                    label: '手机号码',
                    name: 'phone',
                    width: 50,
                    align: 'center'

                }, {
                    label: '发送次数',
                    name: 'count',
                    width: 70,
                    align: 'center'

                },

            ],

            prmNames: {
                page: "page", // 表示请求页码的参数名称
                rows: "rows", // 表示请求行数的参数名称
                sort: "sort",

            },

            ajaxGridOptions: {
                headers: {
                    "LoginToken": loginToken,
                    "Access-Control-Allow-Headers": "LoginToken",
                }
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
                    //toppager: true,
                     loadComplete : function() {
                        var table = this;
                        setTimeout(function(){
                           /* styleCheckbox(table);
                            
                            updateActionIcons(table);*/
                            updatePagerIcons(table);
                           /* enableTooltips(table);*/
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
            
    　
    }

     
});
