var usertype=0
function verify(id) {
    var rowData = $("#jqGrid").jqGrid('getRowData',id);
	var uStatus = rowData.status
    var html=""
    html+='    <table class="checktable">'
    html+='        <tbody>'
    html+='            <tr>'
    html+='                <td >用户名:</td>'
    html+='                <td class="userinfo">'+rowData.username+'</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>公司:</td>'
    html+='                <td class="userinfo">'+rowData.company+'</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>备注:</td>'
    html+='                <td><textarea id="remark" cols="20" rows="2" placeholder="审核备注信息.." style="border: #ccc 1px solid">'+rowData.remark+'</textarea></td>'
    html+='            </tr>'
    html+='        </tbody>'
    html+='    </table>'
  
    var html2=""
    if(uStatus=="通过"){
		html2='<div style="text-align:center" class="check"><label>&nbsp&nbsp&nbsp<input name="isPass" type="radio" id="pass" checked=true/>通过&nbsp&nbsp&nbsp</label><label><input name="isPass" type="radio"  id="notPass"/>不通过</label></div>'
	}else{
		html2='<div style="text-align:center" class="check"><label>&nbsp&nbsp&nbsp<input name="isPass" type="radio" id="pass"/>通过&nbsp&nbsp&nbsp</label><label><input name="isPass" type="radio"  id="notPass"  checked=true/>不通过</label></div>'
	}
    layer.open({
    type: 1,
    skin: 0,
    area: ['300px', '260px'], //宽高
    shade: 0.2,
    title: [
            '审核',
            'background-color:#0b6cbc; color:#fff;'
        ],
    content: html+html2
  ,btn: ['确定']
  ,yes: function(index, layero){
    var isPass = $("[name='isPass']").filter(":checked");
    var status
    if(isPass.attr("id")==undefined){
        layer.msg("状态未选择",{icon: 0});
       return
   } else if(isPass.attr("id")=="pass"){
        status=0
   } else{
        status=2
   }
    $.ajax({
        type: 'post',
        url: "/check_users",
        data: JSON.stringify({
                    "username":rowData.username,
                    "remark":$("#remark").val(),
                    "status": status,
            }),
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        success: function(data) {
          jQuery("#jqGrid").trigger('reloadGrid');
          layer.msg("审核成功",{icon: 1});
        },

        error: function(data) {
            layer.msg("审核失败",{icon: 5});
        },
         complete: function(){
            layer.close(index);
        },
    });

  },
    cancel: function(index){
    layer.close(index);
  }
});
}
function permission(id) {
    var rowData = $("#jqGrid").jqGrid('getRowData',id);
    var html=""
    html+='    <table class="checktable" style="font-size: 17px">'
    html+='        <tbody>'
    html+=            '<tr>'
    html+='                <td >用户名:</td>'
    html+='                <td class="username"><span id="username" >'+rowData.username+'</span></td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>公司名称:</td>'
    html+='                <td><span id="company" >'+rowData.company+'</span></td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>当前可用的接口:</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>鉴黄:</td>'
    html+='                <td class="recognize"><label><input  type="checkbox" id="recognize" /></label></td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>拍屏:</td>'
    html+='                <td class="screen"><label><input  type="checkbox"  id="screen"/></label></td>'
    html+='            </tr>'
/*    html+='            <tr>'
    html+='                <td>暴恐:</td>'
    html+='                <td class="violent"><label><input  type="checkbox" id="violent"/></label></td>'
    html+='            </tr>'*/
    html+='        </tbody>'
    html+='    </table>'

    $.ajax({
        type: 'get',
        url: "/get_usergrant",
        data: {
            "username":rowData.username,
        },
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        success: function(data) {
            jQuery("#jqGrid").trigger('reloadGrid');
            if(data.screen==1){
                $("#screen").attr("checked","checked")
            }
 /*           if(data.violence==1){
                $("#violent").attr("checked","checked")}*/
            if(data.porn==1){
                $("#recognize").attr("checked","checked")}
        },
        error: function(data) {

        },

    });

    layer.open({
    type: 1,
    skin: 0,
    area: ['350px', '250px'], //宽高
    shade: 0.2,
    title: [
            '接口权限',
            'background-color:#0b6cbc; color:#fff;'
        ],
    content: html
  ,btn: ['确定']
  ,yes: function(index, layero){
            var isrecognize = $("#recognize").filter(":checked");
            var recognizestatus
            if(isrecognize.attr("id")=="recognize"){
                recognizestatus=1
            } else{
                recognizestatus=0
            }
            var isscreen = $("#screen").filter(":checked");
            var screenstatus
            if(isscreen.attr("id")=="screen"){
                screenstatus=1
            } else{
                screenstatus=0
            }
            var isviolent = $("#violent").filter(":checked");
            var violentstatus
            if(isviolent.attr("id")=="violent"){
                violentstatus=1
            } else{
                violentstatus=0
            }

    $.ajax({
        type: 'PUT',
        url: "/usergrant",
        data: JSON.stringify({
                    "username":rowData.username,
                    "porn":recognizestatus,
                    "screen":screenstatus,
                    "violence":violentstatus,
            }),
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        success: function(data) {
          jQuery("#jqGrid").trigger('reloadGrid');
          layer.msg("授权成功",{icon: 1});
        },

        error: function(data) {
            layer.msg("授权失败",{icon: 5});
        },
         complete: function(){
            layer.close(index);
        },
    });
  },
    cancel: function(index){
    layer.close(index);
  }
});
}
function GetApisetting(rowData){

    $.ajax({
        type: 'get',
        url: "/get_user_service",
        data: {
            "username":rowData.username,
        },
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        beforeSend: function () {

        },
        success: function (data) {
            usertype=data.usertype
          if(usertype==0){
            $("#free").attr("checked","checked")
          }else if(usertype==1){
            $("#paybefore").attr("checked","checked")
          }else if(usertype==2){
            $("#payafter").attr("checked","checked")
          }else if(usertype==3){
            $("#agent").attr("checked","checked")
          }
          setApiData(data)
        },
        error: function (data) {
        },
    });
}

function setApiData(data){
    $("#currest").val(data.overage)
    $("#dayup").val(data.daily_amount)
    $("#dayrest").val(data.freeday)
    $("#currest").val(data.overage)
    $("#monthtotal").val(data.month_total)
    $("#price").val(data.price)
    $("#concurrency").val(data.concurrency)
    setUserType(data.usertype)
}

function apisetting(id) {
    var rowData = $("#jqGrid").jqGrid('getRowData',id);
    var defaults = {
        setapisetUrl: "/set_user_service",
        getapisetUrl: "/get_user_service",
    };
    GetApisetting(rowData)
    var html=""
    html+=' <div class="apisetbox" > ' +
        '<div class="apisetting"><p class="title">用户类型:</p> ' +
        '<label class="usertype"><input name="isFree" type="radio" id="free" value="0" checked="checked"/>免费</label> ' +
        '<label class="usertype"><input name="isFree" type="radio" value="1" id="paybefore"/>预付费</label></div> ' +
        '<div class="apisetting"><p class="title" style="visibility:hidden;">用户类型:</p> ' +
        '<label class="usertype"><input  name="isFree" type="radio" value="2" id="payafter"/>后付费</label> ' +
        '<label class="usertype"><input name="isFree" type="radio" value="3" id="agent"/>代理商</label></div> ' +
        '<div class="apisetting"><p class="title">用户名:</p><span id="username">'+rowData.username+'</span></div> ' +
        '<div class="free"> ' +
        '<div class="apisetting"><p class="title">每天使用上限:</p><input class="p-num" id="dayup"><span> 次</span></div> </div> ' +
        '<div class="payafter" style="display: none"> ' +
        '<div class="apisetting"><p class="title">单价：</p><input class="p-num" id="price"> <span>元/万张</span></div> ' +
        '<div class="apisetting"><p class="title">当月总用量：</p><input class="p-num" id="monthtotal"> <span> 次</span></div> </div> ' +
        '<div class="apisetting currest"  ><p class="title">当月剩余:</p><input class="p-num" id="currest" disabled="disabled"><span> 次</span></div> ' +
        '<div class="free"> ' +
        '<div class="apisetting"><p class="title">免费剩余天数:</p><input class="p-num" id="dayrest"><span> 天</span></div> </div> ' +
        '<div class="agent" style="display: none"> ' +
        '<div class="apisetting"><p class="title">当月总用量：</p><span>无使用限制</span></div> ' +
        '<div class="apisetting"><p class="title">当前剩余：</p><span>无限制</span></div> ' +
        '<div class="apisetting"><p class="title">计价方式：</p><span>阶梯计价</span></div> </div> ' +
        '<div class="payafter" style="display: none"> <div class="apisetting"><p class="title">用户并发限额：</p><input class="p-num" id="concurrency"></div> </div> ' +
        '<div class="paybefore" style="display: none"> ' +
        '<div class="apisetting">此结算方式暂未开放</div> </div> ' +
        '<div class="apisetting payperiod" style="display: none" ><p class="title">结算周期：</p><select id="payperiod"> ' +
        '<option value="月结">月结</option> ' +
        '<option value="季结">季结</option> ' +
        '<option value="半年结">半年结</option> ' +
        '<option value="年结">年结</option>' +
        '</select></div> ' +
        '</div>'
    layer.open({
        type: 1,
        skin: 0,
        area: ['450px', '450px'], //宽高
        shade: 0.2,
        title: [
                '服务配置',
                'background-color:#0b6cbc; color:#fff;'
            ],
        content: html,
        btn: ['确定'],
        yes: function(index, layero){
            var monthTotal = 0
            if(usertype==0) {
                monthTotal=parseInt($("#dayup").val(), 10)*30
            }else if(usertype==1){
                layer.msg("暂未开放")
                return
            }else if(usertype==2){
                monthTotal=parseInt($("#monthtotal").val(), 10)
            }

            $.ajax({
                type: 'post',
                url: defaults.setapisetUrl,
                data: JSON.stringify({
                    "username":rowData.username,
                    "usertype":parseInt(usertype,10),
                    "month_total":parseInt(monthTotal,10),
                    "freeday":parseInt($("#dayrest").val(), 10),
                    "price":parseInt($("#price").val(), 10),
                    "concurrency": parseInt($("#concurrency").val(), 10),
                    "period": $("#payperiod").val()
                }),
                headers: {
                    "LoginToken": loginToken,
                    "Access-Control-Allow-Headers": "LoginToken",
                },
                success: function () {
                    jQuery("#jqGrid").trigger('reloadGrid');
                    layer.msg('保存成功', {
                        icon: 1
                    });
                    layer.close(index);
                },
                error: function (data) {
                    layer.msg('保存失败', {
                        icon: 5
                    });
                },
            });
        },
        cancel: function(index){
            layer.close(index);
        }
    });
$('input[name="isFree"]').change(function () {
    usertype = ($(this).val());
    setUserType(usertype)
});
}

function setUserType(usertype){
    if(usertype==0){
        $(".payafter").hide()
        $(".free").show()
        $(".agent").hide()
        $(".payperiod").hide()
        $(".currest").show()
        $(".paybefore").hide()
    } else if(usertype==2){
        $(".payafter").show()
        $(".free").hide()
        $(".agent").hide()
        $(".payperiod").show()
        $(".currest").show()
        $(".paybefore").hide()
    } else if(usertype==3){
        $(".payafter").hide()
        $(".free").hide()
        $(".agent").show()
        $(".payperiod").show()
        $(".currest").hide()
        $(".paybefore").hide()
    } else if(usertype==1){
        $(".payafter").hide()
        $(".free").hide()
        $(".agent").hide()
        $(".payperiod").hide()
        $(".currest").hide()
        $(".paybefore").show()
    }
}

function modify(id) {
    var rowData = $("#jqGrid").jqGrid('getRowData',id);
    var phone =rowData.phonenum;
	var utype =rowData.type;	
    var html=""
    html+='    <table class="checktable">'
    html+='        <tbody>'
    html+='            <tr>'
    html+='                <td >用户名:</td>'
    html+='                <td class="userinfo">'+rowData.username+'</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>当前类型:</td>'
    html+='                <td class="userinfo">'+rowData.type+'</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td >联系电话:</td>'
    html+='                <td class="phonenum userinfo"><input id="phonenum" value="'+(phone.substring(2))+'" style="text-indent: 5px"></td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>联系邮箱:</td>'
    html+='                <td class="email userinfo"><input id="email" value="'+rowData.email+'" style="text-indent: 5px"></td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>公司名称:</td>'
    html+='                <td class="userinfo"><input id="company" value="'+rowData.company+'" ></td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>联系人:</td>'
    html+='                <td class="userinfo"><input id="name"  value="'+rowData.name+'"></td>'
    html+='            </tr>'
    html+='        </tbody>'
    html+='    </table>'

    layer.open({
	
    type: 1,
    skin: 0,
    area: ['380px', '300px'], //宽高
    shade: 0.2,
    title: [
            '修改',
            'background-color:#0b6cbc; color:#fff;'
        ],
    content: html,
    btn: ['确定'],
	
    yes: function(index, layero){

    $.ajax({
        type: 'put',
        url: "/userinfo",
        data: JSON.stringify({
                    "username":rowData.username,
                    "phone":'86'+$("#phonenum").val(),
                    "email":$("#email").val(),
                    "company":$("#company").val(),
                    "name": $("#name").val(),
            }),
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        success: function(data) {
          jQuery("#jqGrid").trigger('reloadGrid');
          layer.msg("修改成功",{icon: 1});
        },

        error: function(data) {
            layer.msg("修改失败",{icon: 5});
        },
         complete: function(){
            layer.close(index);
        },
    });

  },
    cancel: function(index){
    layer.close(index);
  }
});
}
function addUser() {
    var defaults = {
        addUserUrl: "/add_users",
    };
    var html=""
    html+='    <table class="checktable" cellspacing="50%" cellpadding="10">'
    html+='        <tbody>'
    html+='            <tr>'
    html+='                <td >用户名:</td>'
    html+='                <td ><input class="p-num" id="username" style="margin-left:10px;width:200px"></td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>密码:</td>'
    html+='                <td><input type="password" class="p-num" id="password" style="margin-left:10px;width:200px"></textarea></td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>确认密码:</td>'
    html+='                <td><input type="password" class="p-num" id="repassword" style="margin-left:10px;width:200px"></textarea></td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>电话号码:</td>'
    html+='                <td><input class="p-num" id="phone"style="margin-left:10px;width:200px"></td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>邮箱:</td>'
    html+='                <td><input class="p-num" id="email" style="margin-left:10px;width:200px"></td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>名字:</td>'
    html+='                <td><input class="p-num" id="name" style="margin-left:10px;width:200px"></td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>公司:</td>'
    html+='                <td><input class="p-num" id="company" style="margin-left:10px;width:200px"></td>'
    html+='            </tr>'
    html+='        </tbody>'
    html+='    </table>'
    layer.open({
    type: 1,
    skin: 0,
    area:  ['360px', '420px'], //宽高
    shade: 0.2,
    title: [
            '添加用户',
            'background-color:#0b6cbc; color:#fff;'
        ],
    content: html
  ,btn: ['确定']
  ,yes: function(index, layero){
        if($("#password").val()!=$("#repassword").val()){
            layer.msg("两次密码不一致")
            return
        }
            $.ajax({
                type: 'post',
                url: defaults.addUserUrl,
                data: JSON.stringify({
                    "username":$("#username").val(),
                    "password": $("#password").val(),
                    "phone":"86"+$("#phone").val(),
                    "email":$("#email").val(),
                    "name":$("#name").val(),
                    "company":$("#company").val(),
                }),
                headers: {
                    "LoginToken": loginToken,
                    "Access-Control-Allow-Headers": "LoginToken",
                },
                success: function () {

                    layer.msg('保存成功', {
                        icon: 1
                    });
                    $("#jqGrid").jqGrid().trigger("reloadGrid");
                    layer.close(index);
                },
                error: function (data) {
                    layer.msg('保存失败,请检查信息是否完整，或者有误', {
                        icon: 5
                    });
                    var status = {
                        702:"用户已存在",
                        603:"用户名格式不正确",
                        601:"用户信息不完整或者格式不正确",
                    };
                    layer.msg(status[data.responseJSON.code] ,{
                        icon: 5,
                    });
                },
            });
  },
    cancel: function(index){
    layer.close(index);
  }
});
}
function delUser(id) {
    var rowData = $("#jqGrid").jqGrid('getRowData',id);
    var defaults = {
        delUserUrl: "/delete_users",
    };
    layer.confirm('是否确认删除用户？', {
        btn: ['确定','取消'] //按钮
    }, function(){
        $.ajax({
            type: 'delete',
            url: defaults.delUserUrl,
            data: JSON.stringify({
                "username":rowData.username,
            }),
            headers: {
                "LoginToken": loginToken,
                "Access-Control-Allow-Headers": "LoginToken",
            },
            success: function () {
                layer.msg('用户已删除', {icon: 1});
                $("#jqGrid").jqGrid().trigger("reloadGrid");
                layer.close();
            },
            error: function (data) {
                layer.msg('用户删除失败', {
                    icon: 5
                });
            },
        });
    })
}

function statistics(id) {
    var rowData = $("#jqGrid").jqGrid('getRowData',id);
    window.location.href="/datastat?username="+rowData.username
}

var type = {
    0:'免费',
    1:'预付费',
    2:'后付费',
    3:'代理商',

}
$(document).ready(function() {
    $(function() {
        pageInit();
    });

    function pageInit() {
        $("#jqGrid").jqGrid({
            url: "/paging_users",
            mtype: "GET",
            colModel: [{
                    label: '用户名',
                    name: 'username',
                    align: 'center',
                },
                {
                    label: '手机号码',
                    name: 'phonenum',
                    align: 'center',
                }, {
                    label: '邮箱',
                    name: 'email',
                    align: 'center'
                }, {
                    label: '公司',
                    name: 'company',
                    align: 'center',
                },{
                    label: '联系人',
                    name: 'name',
                    align: 'center',
                }, {
                    label: '注册时间',
                    name: 'created_at',
                    align: 'center',
                    formatter:'date',
                    formatoptions:{srcformat:'Y-m-d H:i',newformat:'Y-m-d H:i'}

                },{
                    label: '状态',
                    name: 'status',
                    align: 'center',
                },{
                    label: '用户类型',
                    name: 'type',
                    align: 'center',
                    hidden:true,
                    formatter: function (value, grid, rows, state) {
                     return type[value]
                 },
                },{
                    label: '备注',
                    name: 'remark',
                    align: 'center',
                },{  
                    name: '操作', index: '',/* fixed: true, sortable: false, resize: false,*/
                    align: 'center',
                    width:330,
                    //formatter:'actions',  
                    formatter: function (value, grid, rows, state) { 
                        return "<a style=\"color:#f60;cursor:pointer\" onclick=\"verify(" + grid.rowId + ")\">审核</a>&nbsp;" +
                            "<a style=\"color:#f60;cursor:pointer\" onclick=\"modify(" + grid.rowId + ")\">修改</a>&nbsp;" +
                            "<a style=\"color:#f60;cursor:pointer\" onclick=\"statistics(" + grid.rowId + ")\">查看统计</a>&nbsp;" +
                            "<a style=\"color:#f60;cursor:pointer\" onclick=\"permission(" + grid.rowId + ")\">接口权限</a>&nbsp;"+
                            "<a style=\"color:#f60;cursor:pointer\" onclick=\"apisetting(" + grid.rowId + ")\">服务配置</a>&nbsp;"
                    }
                }
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
            beforeProcessing: function (data) {
            data.result = _.map(data.result, function (v) {
                switch (v.status){
                    case 0:
                        v.status="通过";
                    break;
                    case 1:
                        v.status="待审核"
                    break;
                    case 2:
                        v.status="未通过"
                    break;
                 }
                return v
            })
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
            add: true,
            del: true,
            delicon : 'icon-trash red',
            search: false,
            searchicon: 'icon-search orange',
            addicon: 'fa fa-plus orange',
            refresh: true,
            refreshicon: 'icon-refresh green',
            view: true,
            viewicon: 'icon-zoom-in grey',
            addfunc :function(){
                addUser()
            },
           delfunc :function(id){
                delUser(id)
            },

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
