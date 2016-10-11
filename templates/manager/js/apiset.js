/**
 * Created by huang on 2016/8/26.
 */

var defaults = {
    setapisetUrl: "/set_interface_config",
    getapisetUrl: "/get_interface_config",
    getemailsetUrl: "/get_all_notice_email",
    setemailsetUrl: "/add_notice_email",
};


GetApisetting()

GetEmailsetting()

function GetApisetting() {
    $.ajax({
        type: 'get',
        url: defaults.getapisetUrl,
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        beforeSend: function () {

        },
        success: function (data) {
                $("#price").val(data.default_price),
                $("#monthtotal").val(data.default_month_total),
                $("#monthtogethertoatl").val(data.default_concurrency),
                $("#outtime").val(data.validity_time),
                $("#daytotal").val(data.free_day_total),
                $("#daytogethertotal").val(data.free_concurrency)
        },
        complete: function () {

        },
        error: function (data) {
        },
    });
}

function GetEmailsetting() {
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
            $("#email").val(data.email)
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

$("#saveemail").click(function () {
    $.ajax({
        type: 'post',
        url: defaults.setemailsetUrl,
        data: JSON.stringify({
            "email": $("#email").val(),
        }),
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        success: function () {
            layer.msg('保存成功', {
                icon: 1
            });
            GetApisetting()
        },
        error: function (data) {
            layer.msg('保存失败', {
                icon: 5
            });
        },
    });
})

$("#saveapisetting").click(function () {
    $.ajax({
        type: 'post',
        url: defaults.setapisetUrl,
        data: JSON.stringify({
            "default_month_total": parseInt($("#monthtotal").val(), 10),
            "default_price": parseInt($("#price").val(), 10),
            "default_concurrency": parseInt($("#monthtogethertoatl").val(), 10),
            "validity_time": parseInt($("#outtime").val(), 10),
            "free_day_total": parseInt($("#daytotal").val(), 10),
            "free_concurrency": parseInt($("#daytogethertotal").val(), 10),
        }),
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        success: function () {
            layer.msg('保存成功', {
                icon: 1
            });
            GetApisetting()
        },
        error: function (data) {
            layer.msg('保存失败', {
                icon: 5
            });
        },
    });
})

