/**
 * Created by huang on 2016/9/28.
 */

function verify(id) {
    var rowData = $("#jqGrid").jqGrid('getRowData', id);
    var uStatus = rowData.status
    var html = ""
    html += '    <table class="checktable">'
    html += '        <tbody>'
    html += '            <tr>'
    html += '                <td >用户名:</td>'
    html += '                <td class="userinfo">' + rowData.username + '</td>'
    html += '            </tr>'
    html += '            <tr>'
    html += '                <td >联系人:</td>'
    html += '                <td class="userinfo">' + rowData.contact_name + '</td>'
    html += '            </tr>'
    html += '            <tr>'
    html += '                <td >联系电话:</td>'
    html += '                <td class="userinfo">' + rowData.phone + '</td>'
    html += '            </tr>'
    html += '            <tr>'
    html += '                <td>预计用量:</td>'
    html += '                <td class="userinfo">' + rowData.use_total + '</td>'
    html += '            </tr>'
    html += '            <tr>'
    html += '                <td>处理意见:</td>'
    html += '                <td><textarea id="remark" cols="20" rows="3" placeholder="请输入电话回访后的处理意见" style="border: #ccc 1px solid; text-indent: 0;">' + rowData.remark + '</textarea></td>'
    html += '            </tr>'
    html += '        </tbody>'
    html += '    </table>' +
        '<style>.layui-layer-btn .layui-layer-btn0 {border-color: #4898d5;background-color: #2e8ded;color: #fff;display: block;margin: auto;width:125px;text-align: center;} </style>'

    layer.open({
        type: 1,
        skin: 0,
        area: ['400px', '300px'], //宽高
        shade: 0.2,
        title: [
            '套餐升级处理',
            'background-color:#0b6cbc; color:#fff;'
        ],
        content: html
        , btn: ['保存处理结果']
        , yes: function (index, layero) {

            $.ajax({
                type: 'put',
                url: "/packageapply",
                data: JSON.stringify({
                    "username": rowData.username,
                    "remark": $("#remark").val(),
                    package_apply_id: rowData.package_apply_id
                }),
                headers: {
                    "LoginToken": loginToken,
                    "Access-Control-Allow-Headers": "LoginToken",
                },
                success: function (data) {
                    jQuery("#jqGrid").trigger('reloadGrid');
                    layer.msg("处理成功", {icon: 1});
                },

                error: function (data) {
                    var status = {
                        613:"用户的套餐没有找到",
                        701:"用户没有找到",
                    };
                    layer.msg("处理失败", {icon: 5});
                    layer.msg(status[data.responseJSON.code] ,{
                        icon: 5,
                    });
                },
                complete: function () {
                    layer.close(index);
                },
            });

        },
        cancel: function (index) {
            layer.close(index);
        }
    });
}

function delUser(id) {
    var rowData = $("#jqGrid").jqGrid('getRowData', id);
    var defaults = {
        delUserUrl: "/packageapply",
    };
    layer.confirm('是否确认删除用户此用户的套餐升级申请？', {
        btn: ['确定', '取消'] //按钮
    }, function () {
        $.ajax({
            type: 'delete',
            url: defaults.delUserUrl,
            data: JSON.stringify({
                "username": rowData.username,
                package_apply_id: rowData.package_apply_id
            }),
            headers: {
                "LoginToken": loginToken,
                "Access-Control-Allow-Headers": "LoginToken",
            },
            success: function () {
                layer.msg('用户的套餐升级申请已删除', {icon: 1});
                $("#jqGrid").jqGrid().trigger("reloadGrid");
                layer.close();
            },
            error: function (data) {
                layer.msg('用户的套餐升级申请失败', {
                    icon: 5
                });
                var status = {
                    613:"用户的套餐没有找到",
                    701:"用户没有找到",
                };
                layer.msg(status[data.responseJSON.code] ,{
                    icon: 5,
                });
            },
        });
    })
}

$(document).ready(function () {
    $(function () {
        pageInit();
    });

    function pageInit() {
        $("#jqGrid").jqGrid({
            url: "/all_packageapply",
            mtype: "GET",
            colModel: [{
                label: '用户名',
                name: 'username',
                align: 'center',
                width: '90'
            }, {
                label: 'package_apply_id',
                name: 'package_apply_id',
                align: 'center',
                width: '90',
                hidden:true
            }, {
                label: '联系人',
                name: 'contact_name',
                align: 'center',
            }, {
                label: '联系电话',
                name: 'phone',
                align: 'center',
            }, {
                label: '预计用量',
                name: 'use_total',
                align: 'center',
            }, {
                label: '申请时间',
                name: 'CreatedAt',
                align: 'center',
                formatter: 'date',
                formatoptions: {srcformat: 'Y-m-d H:i', newformat: 'Y-m-d H:i'}

            }, {
                label: '处理意见',
                name: 'remark',
                align: 'center',
            }, {
                label: '状态码',
                name: 'status',
                align: 'center',
                hidden: true
            }, {
                name: '操作', index: '',
                align: 'center',
                width: '90',
                formatter: function (value, grid, rows, state) {
                    return "<a style=\"color:#f60;cursor:pointer\" onclick=\"verify(" + grid.rowId + ")\">处理</a>&nbsp;" +
                        "<a style=\"color:#f60;cursor:pointer\" onclick=\"delUser(" + grid.rowId + ")\">删除</a>"
                }

            },
            ],
            prmNames: {
                page: "page", // 表示请求页码的参数名称
                rows: "rows", // 表示请求行数的参数名称
                sort: "sort",
                search: "search",
            },
            postData: {
                status: "0",
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
        url: "/all_packageapply",
        mtype: "GET",
        colModel: [{
            label: '用户名',
            name: 'username',
            align: 'center',
            width: '90'
        },  {
            label: 'package_apply_id',
            name: 'package_apply_id',
            align: 'center',
            width: '90',
            hidden:true
        },{
            label: '联系人',
            name: 'contact_name',
            align: 'center',
        }, {
            label: '联系电话',
            name: 'phone',
            align: 'center',
        }, {
            label: '预计用量',
            name: 'use_total',
            align: 'center',
        }, {
            label: '申请时间',
            name: 'CreatedAt',
            align: 'center',
            formatter: 'date',
            formatoptions: {srcformat: 'Y-m-d H:i', newformat: 'Y-m-d H:i'}

        }, {
            label: '处理意见',
            name: 'remark',
            align: 'center',
        }, {
            label: '状态码',
            name: 'status',
            align: 'center',
            hidden: true
        }, {
            name: '操作', index: '',
            align: 'center',
            width: '90',
            formatter: function (value, grid, rows, state) {
                return "<a style=\"color:#f60;cursor:pointer\" onclick=\"verify(" + grid.rowId + ")\">处理</a>&nbsp;" +
                    "<a style=\"color:#f60;cursor:pointer\" onclick=\"delUser(" + grid.rowId + ")\">删除</a>"
            }

        },
        ],
        prmNames: {
            page: "page", // 表示请求页码的参数名称
            rows: "rows", // 表示请求行数的参数名称
            sort: "sort",
            search: "search",
        },
        postData: {
            status: "0",
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
        url: "/all_packageapply",
        mtype: "GET",
        colModel: [{
            label: '用户名',
            name: 'username',
            align: 'center',
            width: '90'
        },  {
            label: 'package_apply_id',
            name: 'package_apply_id',
            align: 'center',
            width: '90',
            hidden:true
        },{
            label: '联系人',
            name: 'contact_name',
            align: 'center',
        }, {
            label: '联系电话',
            name: 'phone',
            align: 'center',
        }, {
            label: '预计用量',
            name: 'use_total',
            align: 'center',
        }, {
            label: '申请时间',
            name: 'CreatedAt',
            align: 'center',
            formatter: 'date',
            formatoptions: {srcformat: 'Y-m-d H:i', newformat: 'Y-m-d H:i'}
        }, {
            label: '处理意见',
            name: 'remark',
            align: 'center',
        }, {
            label: '状态码',
            name: 'status',
            align: 'center',
            hidden: true
        }, {
            name: '操作', index: '',
            align: 'center',
            width: '90',
            formatter: function (value, grid, rows, state) {
                return "<a style=\"color:#f60;cursor:pointer\" onclick=\"delUser(" + grid.rowId + ")\">删除</a>"
            }

        },
        ],
        prmNames: {
            page: "page", // 表示请求页码的参数名称
            rows: "rows", // 表示请求行数的参数名称
            sort: "sort",
            search: "search",
        },
        postData: {
            status: "1",
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





