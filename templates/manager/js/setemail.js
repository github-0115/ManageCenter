/**
 * Created by huang on 2016/8/31.
 */
var defaults = {
    getemailsetUrl: "/get_all_notice_email",
    setemailsetUrl: "/notice_email",
};

GetEmailsetting()


function GetEmailsetting(){
    $.ajax({
        type: 'get',
        url: defaults.getemailsetUrl,
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        beforeSend: function () {

        },
        success: function (data) {
			var email=""
            data.notice_email.forEach(function(item,idx){
				email+=item.Email+';\n'
            })
			 $(".email").text(email)

        },
        complete: function () {

        },
        error: function (data) {
        },
    });
}

/*$("#addemail").click(function() {

 $.ajax({
 type: 'post',
 url: defaults.setemailsetUrl,
 data: JSON.stringify({
 "email": parseInt($("#email").val(), 10),
 }),
 headers: {
 "LoginToken": loginToken,
 "Access-Control-Allow-Headers": "LoginToken",
 },
 success: function () {
 layer.msg('添加成功', {
 icon: 1
 });
 var emailTable = $('#emailbox');
 var row = $("<tr></tr>");
 var emailHtml = '<td ><input class="email"  placeholder="请输入您要配置的通知邮箱"> </td>';
 row.append(emailHtml);
 emailTable.append(row);

 GetEmailsetting()
 },
 error: function (data) {
 layer.msg('添加失败', {
 icon: 5
 });
 },
 });
 })*/

$("#saveemail").click(function() {
    $.ajax({
        type: 'post',
        url: defaults.setemailsetUrl,
        data: JSON.stringify({
            "email":$(".email").val(),
        }),
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        success: function () {
            layer.msg('保存成功', {
                icon: 1
            });
            GetEmailsetting()
        },
        error: function (data) {
            console.info($(".email").val())
            layer.msg('保存失败', {
                icon: 5
            });
        },
    });
})


