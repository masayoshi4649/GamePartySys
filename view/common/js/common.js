// restpost() ... common REST POST fetch
// restpost("http://192.168.1.42:8080/booking/api/search", {}, {})
// .then((data) => {
//     console.log(data);
//     // DataUpdate
//     DO SOMETHING
// })
// .catch((err) => {
//     console.log(err);
// });
async function restpost(url, header, body) {
    const response = await fetch(url, {
        method: "POST",
        headers: header,
        body: JSON.stringify(body)
    });
    const json = response.json();
    if (response.status == 200) {
        return Promise.resolve(json);
    } else {
        return Promise.reject(json.error);
    }
}


