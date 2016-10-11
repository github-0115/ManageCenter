var checkValue = "EMS邮政快递";
var defaults = {
    searchUrl:"/stat_search_users",
};

$("#LogisticsName").change(function(){
    checkValue=$("#LogisticsName").val();
})

function verify(id) {
    var rowData = $("#jqGrid").jqGrid('getRowData',id);
    var html=""
    html+=' <div class="mskeLayBg"></div>'
    html+=' <div class="mskelayBox">'
    html+='     <div class="mske_html">'
    html+='     </div>'
    html+='     <img class="mskeClaose" src="images/mke_close.png" width="27" height="27" />'
    html+='</div>'
    html+='<div class="msKeimageBox">'
    html+='    <table class="checktable">'
    html+='        <tbody>'
    html+='            <tr>'
    html+='                <td >公司名称:</td>'
    html+='                <td class="userinfo">'+rowData.Username+'</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>发票类型:</td>'
    html+='                <td class="userinfo">'+rowData.Type+'</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>发票抬头:</td>'
    html+='                <td class="userinfo">'+rowData.Title+'</td>'
    html+='            </tr>'
    if(rowData.Type=='企业增值税专用发票'){
    html+='            <tr>'
    html+='                <td >税务登记号:</td>'
    html+='                <td class="userinfo">'+rowData.RegisterNum+'</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>基本开户银行名称:</td>'
    html+='                <td class="userinfo">'+rowData.BankName+'</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>基本开户账号:</td>'
    html+='                <td class="userinfo">'+rowData.BankAccount+'</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>注册场所地址:</td>'
    html+='                <td class="userinfo">'+rowData.BankAddress+'</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>注册固定电话:</td>'
    html+='                <td class="userinfo">'+rowData.insTelNum+'</td>'
    html+='            </tr>'
    html+='            <tr class="photo" >'
    html+='                <td>营业执照复印件:</td>'
    html+='                <td class="userinfo"><img alt="营业执照复印件" src='+rowData.RegisterPhoto+' height="100" width="200" /></td>'
    html+='            </tr>'
    html+='            <tr class="photo">'
    html+='                <td>税务登记复印件:</td>'
    html+='                <td class="userinfo"><img alt="税务登记复印件." src='+rowData.TaxPhoto+' height="100" width="200" /></td>'
    html+='            </tr>'
    html+='            <tr class="photo" >'
    html+='                <td>纳税人资格认证复印件:</td>'
    html+='                <td class="userinfo"><img alt="纳税人资格认证复印件" src='+rowData.TaxpayerPhoto+' height="100" width="200" /></td>'
    html+='            </tr>'
    }
    html+='            <tr>'
    html+='                <td>开票金额:</td>'
    html+='                <td class="userinfo">'+rowData.Money+'</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>收件地址:</td>'
    html+='                <td class="userinfo">'+rowData.Address+'</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>邮编:</td>'
    html+='                <td class="userinfo">'+rowData.ZipCode+'</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>电话:</td>'
    html+='                <td class="userinfo">'+rowData.PhoneNum+'</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>收件人:</td>'   
    html+='                <td class="userinfo">'+rowData.Name+'</td>'
    html+='            </tr>'	
    html+='            <tr>'
    html+='                <td>邮寄信息:</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>物流名称:</td>'
    html+='                <td><select id="LogisticsName"><option value="EMS邮政快递">EMS邮政快递</option><option value="顺丰">顺丰</option></select></td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>快递单号:</td>'
    html+='                <td><input placeholder="请输入快递单号" id="TrackNum"></td>'
    html+='            </tr>'
    html+='        </tbody>'
    html+='    </table>'
    html+='</div>'
    html+='<div class="mskeTogBtn"></div>'
    html+='   <script src="/js/invoices.js"></script>'

    layer.open({
    type: 1,
    skin: 0,
    area: ['600px', '600px'], //宽高
    shade: 0.2,
    title: [
            '开票审核',
            'background-color:#0b6cbc; color:#fff;'
        ],
    content: html,
    end:function(){
        jQuery("#jqGrid").trigger('reloadGrid');
    },
  btn: ['确定']
  ,yes: function(index, layero){
    var status=1
        $.ajax({
        type: 'put',
        url: "/invoiceapply",
        data: JSON.stringify({
                    "username":rowData.Username,
                    "status": status,
                     track_num:$('#TrackNum').val(),
                     logistics_name:checkValue,
                     apply_id:rowData.ApplyId
            }),
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },

        success: function(data) {
          jQuery("#jqGrid").trigger('reloadGrid');
          layer.msg("开票成功",{icon: 1});
        },

        error: function(data) {
            layer.msg("开票失败",{icon: 5});
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
//选项卡滑动切换通用
jQuery(function(){jQuery(".hoverTag .chgBtn").hover(function(){jQuery(this).parent().find(".chgBtn").removeClass("chgCutBtn");jQuery(this).addClass("chgCutBtn");var cutNum=jQuery(this).parent().find(".chgBtn").index(this);jQuery(this).parents(".hoverTag").find(".chgCon").hide();jQuery(this).parents(".hoverTag").find(".chgCon").eq(cutNum).show();})})

//选项卡点击切换通用
jQuery(function(){jQuery(".clickTag .chgBtn").click(function(){jQuery(this).parent().find(".chgBtn").removeClass("chgCutBtn");jQuery(this).addClass("chgCutBtn");var cutNum=jQuery(this).parent().find(".chgBtn").index(this);jQuery(this).parents(".clickTag").find(".chgCon").hide();jQuery(this).parents(".clickTag").find(".chgCon").eq(cutNum).show();})})

//图库弹出层
$(".mskeLayBg").height($(document).height());
$(".mskeClaose").click(function(){$(".mskeLayBg,.mskelayBox").hide()});
$(".msKeimageBox img").click(function(){
   var img=$("<img/>").attr("src",$(this).context.src)
   var max_height=500
    $(".mske_html").html(img);
    $(".mske_html img").css("max-height",max_height+"px")

    var height=img[0].height
    var width=img[0].width
    if(height>max_height){
        height=max_height+15
        width=(img[0].width*(max_height/img[0].height)+15).toFixed(0)
    }
    $(".mskelayBox").css("height",height+"px")
    $(".mskelayBox").css("width",width+"px")
    var outerHeight=-($(".mskelayBox").outerHeight()/2)
    var outerWidth=-($(".mskelayBox").outerWidth()/2)
    $(".mskelayBox").css("margin-top",outerHeight)
    $(".mskelayBox").css("margin-left",outerWidth)
    $(".mskeLayBg").show();$(".mskelayBox").fadeIn(300)
});
//屏蔽页面错误
jQuery(window).error(function(){
  return true;
});
jQuery("img").error(function(){
  $(this).hide();
});

function detail(id) {
    var rowData = $("#jqGrid").jqGrid('getRowData',id);
    $.ajax({
        type: 'GET',
        url: "/invoiceapply",
        data: {
            "user_name":rowData.Username,
            apply_id:rowData.ApplyId
        },
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        success: function(data) {
        },

        error: function(data) {
            layer.msg("查看详情失败",{icon: 5});
        },
        complete: function(){
            layer.close();
        },
    });
    var html=""
    html+=' <div class="mskeLayBg"></div>'
    html+=' <div class="mskelayBox">'
    html+='     <div class="mske_html">'
    html+='     </div>'
    html+='     <img class="mskeClaose" src="images/mke_close.png" width="27" height="27" />'
    html+=' </div>'
    html+='<div class="msKeimageBox">'
    html+='    <table class="checktable">'
    html+='        <tbody>'
    html+='            <tr>'
    html+='                <td >公司名称:</td>'
    html+='                <td class="userinfo">'+rowData.Username+'</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>发票类型:</td>'
    html+='                <td class="userinfo">'+rowData.Type+'</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>发票抬头:</td>'
    html+='                <td class="userinfo">'+rowData.Title+'</td>'
    html+='            </tr>'
    if(rowData.Type=='企业增值税专用发票'){
        html+='            <tr>'
        html+='                <td >税务登记号:</td>'
        html+='                <td class="userinfo">'+rowData.RegisterNum+'</td>'
        html+='            </tr>'
        html+='            <tr>'
        html+='                <td>基本开户银行名称:</td>'
        html+='                <td class="userinfo">'+rowData.BankName+'</td>'
        html+='            </tr>'
        html+='            <tr>'
        html+='                <td>基本开户账号:</td>'
        html+='                <td class="userinfo">'+rowData.BankAccount+'</td>'
        html+='            </tr>'
        html+='            <tr>'
        html+='                <td>注册场所地址:</td>'
        html+='                <td class="userinfo">'+rowData.BankAddress+'</td>'
        html+='            </tr>'
        html+='            <tr>'
        html+='                <td>注册固定电话:</td>'
        html+='                <td class="userinfo">'+rowData.insTelNum+'</td>'
        html+='            </tr>'
        html+='            <tr class="photo" >'
        html+='                <td>营业执照复印件:</td>'
        html+='                <td class="userinfo"><img alt="营业执照复印件" src='+rowData.RegisterPhoto+' height="100" width="200" /></td>'
        html+='            </tr>'
        html+='            <tr class="photo">'
        html+='                <td>税务登记复印件:</td>'
        html+='                <td class="userinfo"><img alt="税务登记复印件." src='+rowData.TaxPhoto+' height="100" width="200" /></td>'
        html+='            </tr>'
        html+='            <tr class="photo" >'
        html+='                <td>纳税人资格认证复印件:</td>'
        html+='                <td class="userinfo"><img alt="纳税人资格认证复印件" src='+rowData.TaxpayerPhoto+' height="100" width="200" /></td>'
        html+='            </tr>'
    }
    html+='            <tr>'
    html+='                <td>开票金额:</td>'
    html+='                <td class="userinfo">'+rowData.Money+'</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>收件地址:</td>'
    html+='                <td class="userinfo">'+rowData.Address+'</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>邮编:</td>'
    html+='                <td class="userinfo">'+rowData.ZipCode+'</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>电话:</td>'
    html+='                <td class="userinfo">'+rowData.PhoneNum+'</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>收件人:</td>'
    html+='                <td class="userinfo">'+rowData.Name+'</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>邮寄信息:</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>物流名称:</td>'
    html+='                <td class="userinfo">'+rowData.LogisticsName+'</td>'
    html+='            </tr>'
    html+='            <tr>'
    html+='                <td>快递单号:</td>'
    html+='                <td class="userinfo">'+rowData.TrackNum+'</td>'
    html+='            </tr>'
    html+='        </tbody>'
    html+='    </table>'
    html+='</div>'
    html+='<div class="mskeTogBtn"></div>'
    html+='   <script src="/js/invoices.js"></script>'
    layer.open({
    type: 1,
    skin: 0,
    area: ['600px', '600px'], //宽高
    shade: 0.2,
    title: [
            '开票详情',
            'background-color:#0b6cbc; color:#fff;'
        ],
    content: html,
    cancel:function(index){
       layer.close(index)
    },
});
}

var type = {
    0:'企业增值税普通发票',
    1:'企业增值税专用发票'
}

$(document).ready(function() {
    var username=""
    var dataPost={}
    var invoiceUrl="/page_invoiceapply"
    function getUrlParam(name) {
        var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
        var r = window.location.search.substr(1).match(reg);  //匹配目标参数
        if (r != null) return unescape(r[2]);
        return null; //返回参数值
    }

    $(function() {
        if (getUrlParam("username") != null) {
            username = getUrlParam("username");
            invoiceUrl="/user_invoiceapply"
            dataPost={
                user_name:username,
            }
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
                window.location.href="/searchuser?url=invoices&text="+$("#autocomplete-search").val()
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
           jQuery("#jqGrid").jqGrid('setGridParam',{
            url: "/user_invoiceapply",
            postData:{
                user_name:username,
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
            url: invoiceUrl,
            mtype: "GET",
            colModel: [{
                    label: '公司名',
                    name: 'Username',
                    align: 'center',
                },{
                    label: '发票id',
                    name: 'InvoiceId',
                    align: 'center',
                    hidden:true
                },{
                    label: '申请id',
                    name: 'ApplyId',
                    align: 'center',
                    hidden:true
                },{
                    label: 'AddressId',
                    name: 'AddressId',
                    align: 'center',
                    hidden:true
                },
                {
                    label: '邮编',
                    name: 'ZipCode',
                    align: 'center',
                    hidden:true
                },
                {
                    label: '税务登记号',
                    name: 'RegisterNum',
                    align: 'center',
                    hidden:true
                },{
                    label: '基本开户银行名称',
                    name: 'BankName',
                    align: 'center',
                    hidden:true
                },{
                    label: '基本开户账号',
                    name: 'BankAccount',
                    align: 'center',
                    hidden:true
                },{
                    label: '注册场所地址',
                    name: 'BankAddress',
                    align: 'center',
                    hidden:true
                },{
                    label: '注册固定电话',
                    name: 'insTelNum',
                    align: 'center',
                    hidden:true
                },{
                    label: '营业执照复印件',
                    name: 'RegisterPhoto',
                    align: 'center',
                    hidden:true
                },
                {
                    label: '税务登记复印件',
                    name: 'TaxPhoto',
                    align: 'center',
                    hidden:true
                },
                {
                    label: '纳税人资格认证复印件',
                    name: 'TaxpayerPhoto',
                    align: 'center',
                    hidden:true
                },{
                    label: '收件人',
                    name: 'Name',
                    align: 'center',
                    hidden:true
                },{
                    label: '发票抬头',
                    name: 'Title',
                    align: 'center',
                    hidden:true
                },{
                    label: '开票金额',
                    name: 'Money',
                    align: 'center',

                }, {
                    label: '发票类型',
                    name: 'Type',
                    align: 'center',
                    formatter: function (value, grid, rows, state) {
                        return type[value]
                    }

                }, {
                    label: '联系电话',
                    name: 'PhoneNum',
                    align: 'center',
                },{
                    label: '联系地址',
                    name: 'Address',
                    align: 'center',
                    width:250,
                },{
                    label: '物流名称',
                    name: 'LogisticsName',
                    align: 'center',
                    hidden:true
                },{
                    label: '快递单号',
                    name: 'TrackNum',
                    align: 'center',
                    hidden:true
                }, {
                    label: '状态',
                    name: 'Status',
                    align: 'center',
                },{
                    name: '操作', index: '',/* fixed: true, sortable: false, resize: false,*/
                    align: 'center',
                    width:200,
                    //formatter:'actions',  
                    formatter: function (value, grid, rows, state) {
                        if(rows.Status=='未处理')
                        return "<a style=\"color:#f60;cursor:pointer\" onclick=\"verify(" + grid.rowId + ")\">审核操作</a>";
                        else{
                            return"<a style=\"color:#428bca;cursor:pointer\" onclick=\"detail(" + grid.rowId + ")\">查看详情</a>";
                          }
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
            postData: dataPost,
            jsonReader: {
                root: "apply",
                page: "page",
                total: "total",
                records: "record",
                repeatitems: false
            },
            rowList: [15, 25, 25],
            viewrecords: true, 
            autowidth:true,
       
            height: 600,
            rowNum: 15,
            datatype: "json",
            pager: "#jqGridPager",
            autosearch: true,
           
            altRows: true,
            beforeProcessing: function (data) {
            data.apply = _.map(data.history.apply, function (v) {
                switch (v.Status){
                    case 0:
                        v.Status="未处理";
                    break;
                    case 1:
                        v.Status="已处理"
                    break;
                 }
                return v
            })
            var apply=[]
            data.history.apply.forEach(function(item,idx){
                var temp={}
                temp={
                    "ApplyId":item.ApplyId ,
                    "Username":item.Username ,
                    "InvoiceId": item.InvoiceId,
                    "AddressId":item.AddressId ,
                    "LogisticsName":item.LogisticsName,
                    "TrackNum":item.TrackNum,
                    "Status":item.Status,
                    "Money":(item.Money/100).toFixed(2)+"元",
                    "Type": data.history.invoice[idx].Type,
                    "Title": data.history.invoice[idx].Title,
                    "RegisterNum":data.history.invoice[idx].RegisterNum ,
                    "BankName":data.history.invoice[idx].BankName ,
					"RegisterPhoto":data.history.invoice[idx].RegisterPhoto ,
					"TaxPhoto":data.history.invoice[idx].TaxPhoto ,
					"TaxpayerPhoto":data.history.invoice[idx].TaxpayerPhoto ,
                    "BankAccount":data.history.invoice[idx].BankAccount ,
                    "BankAddress":data.history.invoice[idx].BankAddress ,
                    "insTelNum":data.history.invoice[idx].TelNum ,
                    "Address": data.history.address[idx].Province+data.history.address[idx].City+data.history.address[idx].District+data.history.address[idx].Address,
                   /* "Address":procity+data.history.address[idx].District+data.history.address[idx].Address,*/
                    "ZipCode":data.history.address[idx].ZipCode ,
                    "PhoneNum":data.history.address[idx].PhoneNum ,
                    "Name": data.history.address[idx].Name,
                    "addrTelNum": data.history.address[idx].TelNum,
                }

                apply.push(temp)
            })
                data.apply = apply
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
});
