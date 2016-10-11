/**
 * Created by huang on 2016/7/20.
 */
$(document).ready(function() {
    var cnResult= {
            p:"色情",
            s:"性感",
            n:"正常",
            screen:"拍屏",
    };
    var defaults = {
        searchUrl:"/stat_search_users",
    };
    var username=""
    var dataPost={}
    var viewUrl="/api/view"
    function getUrlParam(name) {
        var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
        var r = window.location.search.substr(1).match(reg);  //匹配目标参数
        if (r != null) return unescape(r[2]);
        return null; //返回参数值
    }

    $(function() {
        if (getUrlParam("username") == null) {
            viewUrl="/api/view"
        }else{
            username = getUrlParam("username");
            viewUrl="/api/user_view"
            dataPost={
                username:username,
            }
            $("#lblUser").text("当前查看"+username+"的数据")
        }
        pageInit();
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
                window.location.href="/searchuser?url=alldata&text="+$("#autocomplete-search").val()
            },
            error: function (data) {
                //window.location.href="/searchuser "
            },
        });
    });

    $('#autocomplete-search').autocomplete({
       serviceUrl: '/stat_search_users',
        params:{
            page:"1",
            rows:"10",
        },
        ajaxSettings:{
            headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
            },
        },
        onSelect: function (suggestion) {
           username=suggestion.data
           $("#lblUser").text("当前查看"+username+"的数据")
           jQuery("#jqGrid").jqGrid('setGridParam',{
            url: viewUrl,
            postData:{
                page:1,
                rows:20,
                username:username,
            },
        }).trigger("reloadGrid"); 
        },
        transformResult: function(response) {
             response=$.parseJSON(response);
             var result=[]
             response.result.forEach(function(item,idx){
                var resultTemp={}
                resultTemp={
                    "username":item.username,
                    "company":item.company
                }
                result.push(resultTemp)
                 if(item.company==""){
                    resultTemp={
                        "username":item.username,
                        "company":item.username
                    }
                    result.push(resultTemp)
                }
             })
                return {
            suggestions: $.map(result, function(dataItem) {
                return { value: dataItem.company, data: dataItem.username};
            })
        }
        }
    });


    function pageInit() {
         $("#jqGrid").jqGrid({
            url: viewUrl,
            mtype: "GET",
            styleUI: 'Bootstrap',
             colModel: [
                {
                    label: '时间',
                    name: 'created_at',
                    width: 40,
                    align: 'center',
                    formatter: 'created_at',
                    formatoptions: {srcformat: 'Y-m-d H:i', newformat: 'Y-m-d H:i'},
                    cellattr: function(rowId, tv, rawObject, cm, rdata) {
                        return 'id=\'created_at' + rowId + "\'";
                    }

                },
                {
                    label: '分类',
                    name: 'type',
                    width: 40,
                    align: 'center',
                }, {
                    label: '确认量',
                    name: 'type_check',
                    width: 50,
                    align: 'center',
                },
                {
                    label: '复审量',
                    name: 'type_review',
                    width: 50,
                    align: 'center',
                },{
                    label: '总量',
                    name: 'total',
                    width: 50,
                    align: 'center',
                    cellattr: function(rowId, tv, rawObject, cm, rdata) {
                        return 'id=\'total' + rowId + "\'";
                    }
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
            postData: dataPost,
            jsonReader: {
                root: "result",
                page: "page",
                total: "total",
                records: "records",
                repeatitems: false
            },
            viewrecords: false,
            autowidth: true,
            height: "auto",
            rowNum: 20,
            datatype: "json",
            pager: "#jqGridPager",
            autosearch: true,

            altRows: true,
            gridComplete: function() {
                var gridName = "jqGrid";
                Merger(gridName, 'created_at');
                Merger(gridName, 'total');
            },
            beforeProcessing: function (data) {
                 var result=[]
                var toPercent = function (v) {
                    if(isNaN(v))
                        return "0.00%"
                    else
                        return (v * 100).toFixed(2) + "%"
                }
                var rateSummary = function (typ,item) {
                   typ=cnResult[typ]+'复审占总比:'+toPercent(item[typ]/item.total) ;
                }
                data.result.forEach(function (item,idx){
                    ['n','s','p','screen'].forEach(function (typ) {
                        var temp={}
                        if(typ=='screen'){
                            temp={
                                created_at:item.time,
                                type:cnResult[typ],
                                type_check:(item.isScreen).toLocaleString()+'/'+item.notScreen.toLocaleString(),
                                type_review:'-',
                                total:item.total.toLocaleString()+'</br>'+'正常复审占总比:'+toPercent(item['n_review']/item.total)+'</br>'+
                                '性感复审占总比:'+toPercent(item['s_review']/item.total)+'</br>'+'色情复审占总比:'+toPercent(item['p_review']/item.total),
                            }
                        }else{
                             temp={
                                created_at:item.time,
                                type:cnResult[typ],
                                type_check:(item[typ]-item.p_review).toLocaleString(),
                                type_review:item[typ+'_review'].toLocaleString()+"("+toPercent(item[typ+'_review']/item[typ])+")",
                                total:item.total.toLocaleString()+'</br>'+'正常复审占总比:'+toPercent(item['n_review']/item.total)+'</br>'+
                                '性感复审占总比:'+toPercent(item['s_review']/item.total)+'</br>'+'色情复审占总比:'+toPercent(item['p_review']/item.total),
                            }
                        }
                        temp[typ]=item[typ]
                        result.push(temp)
                    })
                })
                data.result=result
            },
            loadComplete : function() {
                var table = this;
                setTimeout(function(){

                    updatePagerIcons(table);
                }, 0);
            },

        });

 function Merger(gridName, CellName) {
            //得到显示到界面的id集合
            var mya = $("#" + gridName + "").getDataIDs();
            //当前显示多少条
            var length = mya.length;
            for (var i = 0; i < length; i++) {
                //从上到下获取一条信息
                var before = $("#" + gridName + "").jqGrid('getRowData', mya[i]);
                //定义合并行数
                var rowSpanTaxCount = 1;
                for (j = i + 1; j <= length; j++) {
                    //和上边的信息对比 如果值一样就合并行数+1 然后设置rowspan 让当前单元格隐藏
                    var end = $("#" + gridName + "").jqGrid('getRowData', mya[j]);
                    if (before[CellName] == end[CellName]) {
                        rowSpanTaxCount++;
                        $("#" + gridName + "").setCell(mya[j], CellName, '', { display: 'none' });
                    } else {
                        rowSpanTaxCount = 1;
                        break;
                    }
                    $("#" + CellName + "" + mya[i] + "").attr("rowspan", rowSpanTaxCount);
                }
            }
        }

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


});
