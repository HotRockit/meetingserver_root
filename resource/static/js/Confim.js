function deleteUser() {
    let sta = confirm("确定删除此帐户吗？");
    if (sta == true) {
        $.ajax({
            url: "/user/deleteSuperUser",
            type: 'POST',
            dataType: 'text',
            data: {username: $("#button2").val()},
        })
            .done(function (data) {
                data = JSON.parse(data)
                if (data.data==1) {
                    // alert("删除成功");
                    window.location.href = "/user/login";
                } else {
                    alert("删除失败");
                }
            })
            .fail(function () {
                console.log("error");
            })
            .always(function () {
                console.log("complete");
            });
    }
}

function changePassword() {
    let sta = confirm("确定修改此密码吗？");
    if (sta == true) {
        $.ajax({
            url: "/user/updatePassword",
            type: 'POST',
            dataType: 'text',
            data: {
                username: $("#button1").val(),
                password: $("#button2").val(),
            },
        })
            .done(function (data) {
                data = JSON.parse(data)
                if (data.data==1) {
                    // alert("修改成功");
                    window.location.href = "/user/login";
                } else {
                    alert("修改失败");
                }
            })
            .fail(function () {
                console.log("error");
            })
            .always(function () {
                console.log("complete");
            });
    }
}


function back() {
    let st = confirm("是否退出?");
    if (st == true) {
        // $.ajax({
        //     url: "/doback.do",
        //     type: 'POST',
        //     dataType: 'text',
        //     // data: {
        //     //     username:$("#backli").attr("value")
        //     // },
        // })
        //     .done(function (data) {
        //         if (data == 0) {
        //             alert("退出成功");
        //             window.location.href = "/user/login";
        //         } else {
        //             alert("退出失败");
        //         }
        //     })
        //     .fail(function () {
        //         console.log("error");
        //     })
        //     .always(function () {
        //         console.log("complete");
        //     });
        window.location.href = "/user/login"
    }
}

function leveljudge() {

    let a = $("#leveljudge").attr("value");
    if (a == 1 || a == 2 || a == 21) {
        alert("你的身份权限不够");
        $("#leveljudge1").hide();
        $("#leveljudge").css("color", "#EE5C42");
    }
}

function update() {
    window.location.href = "/user/userEdit";
}


