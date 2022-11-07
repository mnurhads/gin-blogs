$(document).ready(function () {
    $("register-from").validate({
        rules: {
            username: {
                required: true,
                rangelength: [5, 10]
            },
            password: {
                required: true,
                rangelength: [5, 10]
            },
            repassword: {
                required: true,
                rangelength: [5, 10],
                equalTo: "#register-password"
            }
        },
        messages: {
            username: {
                required: "Username Wajib diisi",
                rangelength: "Panjang username harus 5-10"
            },
            password: {
                required: "Password Wajib Diisi",
                rangelength: "Panjang Password Wajib harus 5-10"
            },
            repassword: {
                required: "Konfirmasi Password Wajib diisi",
                rangelength: "Panjang Konfirmasi wajib 5-10",
                equalTo: "Password dan KOnfirmasi tidak sama"
            }
        },
        submitHandler: function (form) {
            var urlStr = "/register";
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert("data:" + data.message)
                    if (data.code == 1) {
                        setTimeout(function () {
                            window.location.href = "/login"
                        }, 1000)
                    }
                },
                err: function (data, status) {
                    alert("err:" + data.message + ":" + status)
                }
            })
        }
    });


    $("#login-form").validate({
        rules: {
            username: {
                required: true,
                rangelength: [5, 10]
            },
            password: {
                required: true,
                rangelength: [5, 10]
            }
        },
        messages: {
            username: {
                required: "Username wajib diisi",
                rangelength: "Panjang Username harus 5-10"
            },
            password: {
                required: "Password wajib diisi",
                rangelength: "Panjang password harus 5-10"
            }
        },
        submitHandler: function (form) {
            var urlStr = "/login"
            alert("urlStr:" + urlStr)
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert("data:" + data.message + ":" + status)
                    if (data.code == 1) {
                        setTimeout(function () {
                            window.location.href = "/"
                        }, 1000)
                    }
                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)

                }
            });
        }
    });

    $("#write-art-form").validate({
        rules: {
            title: "required",
            tags: "required",
            short: {
                required: true,
                minlength: 2
            },
            content: {
                required: true,
                minlength: 2
            }
        },
        messages: {
            title: "Silakan masukkan judul",
            tags: "Harap masukkan label",
            short: {
                required: "Silakan masukkan profil",
                minlength: "setidaknya dua karakter"
            },
            content: {
                required: "Silakan masukkan isi artikel",
                minlength: "Isi artikel minimal dua karakter"
            }
        },
        submitHandler: function (form) {
            alert("hello")
            var urlStr = "/article/add";
            var artId = $("#write-article-id").val();
            alert("artId:" + artId);
            if (artId > 0) {
                urlStr = "/article/update"
            }
            alert("urlStr:" + urlStr);
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert(":data:" + data.message);
                    setTimeout(function () {
                        window.location.href = "/"
                    }, 1000)
                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)
                }
            });
        }
    });

    $("#album-upload-button").click(function () {
        var filedata = $("#album-upload-file").val();
        if (filedata.length <= 0) {
            alert("Upload melebihi batas!");
            return
        }

        var data = new FormData()
        data.append("upload", $("#album-upload-file")[0].files[0]);
        alert(data)
        var urlStr = "/upload"
        $.ajax({
            url: urlStr,
            type: "post",
            dataType: "json",
            contentType: false,
            data: data,
            processData: false,
            success: function (data, status) {
                alert(":data:" + data.message);
                if (data.code == 1) {
                    setTimeout(function () {
                        window.location.href = "/album"
                    }, 1000)
                }
            },
            error: function (data, status) {
                alert("err:" + data.message + ":" + status)
            }
        })
    })
});