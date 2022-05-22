let table = new Tabulator("#recruitTable", {
    selectable: 1,
    height: "40vh",
    layout: "fitDataStretch",
    columns: [
        { title: "種別", field: "RecruitTypeName", hozAlign: "left", width: 100, },
        { title: "題名", field: "RecruitTitle", hozAlign: "left", width: 150 },
        { title: "場所", field: "RecruitPlace", hozAlign: "left", width: 150, },
        { title: "現在人数", field: "MinJoin", width: 80, hozAlign: "center", },
        { title: "最大人数", field: "MaxJoin", width: 80, hozAlign: "center", },
        {
            title: "締め切り日時", field: "RecruitDue", hozAlign: "center",
            formatter: function (cell) {
                return formatDateObj(cell.getValue())
            },
        },
        {
            title: "🔒",
            field: "JoinPassExist",
            headerTooltip: "パスワード設定あり",
            headerHozAlign: "center",
            formatter: function (cell, formatterParams, onRendered) {
                switch (cell.getValue()) {
                    case true:
                        return "🔒";

                    case false:
                        return null;
                }
            },
            width: 54,
            hozAlign: "center",
        },
        {
            title: "❗",
            field: "RecruitCommentMust",
            headerTooltip: "必読メッセージあり",
            headerHozAlign: "center",
            formatter: function (cell) {
                switch (cell.getValue()) {
                    case true:
                        return "❗";

                    case false:
                        return null;
                }
            },
            width: 54,
            hozAlign: "center",
        },
        {
            title: "求める意気込み", field: "PlayStyle", formatter: "star", formatterParams: {
                stars: 5,
            }
        },
    ],
});

table.on("rowSelectionChanged", function (data, rows) {
    switch (data.length) {
        case 1:
            document.getElementById("details").innerHTML = getDetails(data[0]);
            document.getElementById("creatept").innerHTML = null;
            break;

        default:
            document.getElementById("details").innerHTML = null;
            document.getElementById("creatept").innerHTML = getCreatept();
            break;
    }
});

function getCreatept() {
    return `
    <button class="createptbtn"role="button">
    募集する
    </button>
    `
}

function getDetails(json) {
    console.log(json.DisplayUserIcon)

    return `
    募集種別:${json.RecruitTypeName}
    <br>
    題名:${json.RecruitTitle}
    <br>
    場所:${json.RecruitPlace}
    <br>
    人数:最大:${json.MaxJoin} / 最小:${json.MinJoin}
    <br>
    パスワード:${dispJoinPassExist(json.JoinPassExist)}
    <br>
    必読コメント:${dispRecruitCommentMust(json.RecruitCommentMust)}
    コメント:${json.RecruitComment.String}
    <br>
    求める意気込み:${dispPlayStyle(json.PlayStyle)}
    <br>
    募集ユーザ:${json.DisplayUserName}
    <br>
    <img src="${json.DisplayUserIcon}">
    <br>
    募集締め切り:${formatDateObj(json.RecruitDue)}
    <br>
    募集開始:${formatDateObj(json.CreateTime)}
`
}

function dispJoinPassExist(exist) {
    switch (exist) {
        case true:
            return "あり"

        case false:
            return "なし"
    }
}

function dispRecruitCommentMust(exist) {
    switch (exist) {
        case true:
            return "あり"

        case false:
            return "なし"
    }
}

function dispPlayStyle(value) {
    let stars = "";
    for (let i = 0; i < value; i++) {
        stars = stars.concat("🌟")
    }
    return stars
}

window.onload = () => {
    updateData()
}

function updateData() {
    restpost("http://192.168.1.42:8080/booking/api/search", {}, {})
        .then((data) => {
            console.log(data);
            table.setData(data);
        })
        .catch((err) => {
            console.log(err);
        });
}

function formatDateObj(date) {
    const dateobj = getDateObj(date)
    return dateobj["year"] + "/" + dateobj["month"] + "/" + dateobj["date"] + " " + dateobj["hour"] + ":" + dateobj["minute"] + ":" + dateobj["second"]
}

function getDateObj(date) {
    const dateobj = new Date(date);
    return {
        "year": dateobj.getFullYear(),
        "month": zeroPadding((dateobj.getMonth() + 1), 2),
        "date": zeroPadding(dateobj.getDate(), 2),
        "hour": zeroPadding(dateobj.getHours(), 2),
        "minute": zeroPadding(dateobj.getMinutes(), 2),
        "second": zeroPadding(dateobj.getSeconds(), 2),
        "time": zeroPadding(dateobj.getTime(), 3),
    };
}

function zeroPadding(num, length) {
    return ("00000" + num).slice(-length);
}

