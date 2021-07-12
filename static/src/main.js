fetch("/hello")
  .then((response) => response.json())
  .then((data) => {
    console.log(data);
    document.getElementById("root").innerHTML = `
    <div>
        <p>hello endpoint: ${data.message} </p>
    </div>`;
  });
fetch("/ping")
  .then((response) => response.json())
  .then((data) => {
    console.log(data);
    var rootdiv = document.getElementById("root");
    var newcontent = document.createElement("div");
    newcontent.innerHTML = `<div><p>ping endpoint: ${data.message} </p></div>`;
    rootdiv.appendChild(newcontent.firstChild);
  });
