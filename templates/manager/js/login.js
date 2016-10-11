$(document).ready(function() {
    var defaults = {
        loginrUrl: "/login_token",
    };
    document.onkeydown = function(e){

        var ev = document.all ? window.event : e;

        if(ev.keyCode==13) {
            $("#login").click();
        }

    }
    $("#managername").change(function() {
        var flag = validateName($(this).val())
        if (flag) {
            layer.tips(flag, '#managername', {
                tipsMore: true
            });
        }
    });
    $("#password").change(function() {
        var flag = validatePassword($(this).val())
        if (flag) {
            layer.tips(flag, '#password', {
                tipsMore: true
            });
        }
    });

    function validateName(managername) {
        if (managername == "") {
            return '请输入用户名！';
        }
        return false;
    }

    function validatePassword(password) {     
        if (password == "") {
            return '请输入密码！';
        }
        return false;
    }

    $("#login").click(function() {


        var managername = $("#managername").val();
        var password = $("#password").val();

        var userFlag = validateName(managername);
        var passwordFlag = validatePassword(password);

        if (userFlag) {
            layer.tips(userFlag, '#managername', {
                tipsMore: true
            });
        }

        if (passwordFlag) {
            layer.tips(passwordFlag, '#password', {
                tipsMore: true
            });
        }

        if (!userFlag && !passwordFlag) {

            console.debug(managername);
            console.debug(password);
            $.ajax({
                type: 'post',
                url: defaults.loginrUrl,
                data: JSON.stringify({
                    managername: managername,
                    password: password
                }),
                contentType: "application/json; charset=utf-8",
                dataType: 'json',
                success: function(data) {
                    window.localStorage.setItem("LoginToken", data.token);
                    window.localStorage.setItem("managername", managername);
                    $("#response").html(data);
                    layer.msg('登录成功', {
                        icon: 1
                    });

                    window.location.href = "index";
                },

                error: function(data) {
                    layer.msg('登录失败，请检查您的用户名密码是否正确', {
                        icon: 5
                    });
                },
                beforeSend: function() {
                    $("#response").html("send request...");


                }
            });

        }
        return false;
    });

});
