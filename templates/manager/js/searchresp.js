/**
 * Created by huang on 2016/7/21.
 */
$(document).ready(function() {
    var defaults = {
        searchUrl: "/api/user_view",
    };
    $(function () {
        pageInit();
    });
    function pageInit() {
        $("#jqGrid").jqGrid({
            url: defaults.searchUrl,
            mtype: "GET",
            styleUI: 'Bootstrap',
            colModel: [
                {
                    label: '用户名',
                    name: 'user',
                    width: 30,
                    align: 'center',
                    searchoptions: {sopt: ['eq']}
                },

                {
                    label: '时间',
                    name: 'created_at',
                    width: 30,
                    align: 'center',
                    formatter: 'time',
                    formatoptions: {srcformat: 'Y-m-d H:i', newformat: 'Y-m-d H:i'}
                },
                {
                    label: '总量',
                    name: 'total',
                    width: 30,
                    align: 'center',
                    searchoptions: {sopt: ['eq']}
                }, {
                    label: '正常/比例',
                    name: 'n_rat',
                    width: 40,
                    align: 'center',


                },
                {
                    label: '色情/比例',
                    name: 'p_rat',
                    width: 40,
                    align: 'center',


                }, {
                    label: '性感/比例',
                    name: 's_rat',
                    width: 40,
                    align: 'center'


                }, {
                    label: '正常（review)/占正常比例/总的比例',
                    name: 'n_str',
                    width: 70,
                    align: 'center',
                },
                {
                    label: '色情（review)/占色情比例/总的比例',
                    name: 'p_str',
                    width: 70,
                    align: 'center',
                },
                {
                    label: '性感（review)/占性感比例/总的比例',
                    name: 's_str',
                    width: 70,
                    align: 'center',
                },

            ],

            prmNames: {
                page: "page", // 表示请求页码的参数名称
                rows: "rows", // 表示请求行数的参数名称
                sort: "sort",
                search: "search",
            },

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
                records: "records",
                repeatitems: false
            },
            rowList: [15, 20, 25],
            viewrecords: true,
            autowidth: true,
            height: 500,
            rowNum: 15,
            datatype: "json",
            pager: "#jqGridPager",
            autosearch: true,

            altRows: true,
            beforeProcessing: function (data) {
                var toPercent = function (v) {
                    if(isNaN(v))
                        return "0.00"+"%"
                    return (v * 100).toFixed(2) + "%"
                }
                var reviewSummary = function (v, type) {
                    var review = v[type + '_review']
                    v[type + '_str'] = review + " / " + toPercent(review / v[type]) + " / " + toPercent(v[type + '_review']/ v.total);
                }
                data.result = _.map(data.result, function (v) {
                    ['p', 's','n'].forEach(function (typ) {
                        reviewSummary(v, typ)
                    })
                    return v
                })
                var rateSummary = function (v, type) {
                    v[type + '_rat'] =  v[type]  + " / " + toPercent(v[type] / v.total) ;
                }

                data.result = _.map(data.result, function (v) {
                    ['n','p','s'].forEach(function (typ) {
                        rateSummary(v, typ)
                    })
                    return v
                })

                var company =  $(".company").val();
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
            view: false,
            viewicon: 'icon-zoom-in grey',

        });


    }

    })