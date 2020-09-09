    const emptyInput = "Musisz podać jakiś URL."
    const bloodyRedTheme = "css/red-style.css"
    const darkNightTheme = "css/dark-style.css"
    const themeSwitch = document.getElementById("theme-switch")

    window.onload = function () {
        localStorage.getItem(themeSwitch.id) === "true"
            ? checkThemeSwitchAndUpdateCss(bloodyRedTheme)
            : uncheckThemeSwitchAndUpdateCss(darkNightTheme)

        setTimeout(() => $("#cover").fadeOut(500), 1000);
    }

    window.addEventListener("keydown", function(event) {
        if (event.key === "Enter") {
            event.preventDefault();
            submitForm();
        }
    });

    document.getElementById("submit-button").addEventListener("click", function(event){
        event.preventDefault()
        submitForm()
    });

    document.getElementById("show-tailored-cut").addEventListener("click", function(event){
        let icon = document.getElementById("tailored-cut-icon")

        if (icon.classList.contains("fa-plus-circle")) {
            icon.classList.replace("fa-plus-circle", "fa-minus-circle")

        } else {
            icon.classList.replace("fa-minus-circle", "fa-plus-circle")
        }
    });

    themeSwitch.addEventListener("click", function (event){
        let style = document.getElementById("style-select");

        this.checked === true
            ? setAndStoreTheme(style, bloodyRedTheme, this)
            : setAndStoreTheme(style, darkNightTheme, this)
    })

    function post(url, data, success) {
        let params = typeof data == 'string' ? data : Object.keys(data).map(
        function(k){ return encodeURIComponent(k) + '=' + encodeURIComponent(data[k]) }
        ).join('&');

        let xhr = window.XMLHttpRequest ? new XMLHttpRequest() : new ActiveXObject("Microsoft.XMLHTTP");
        xhr.open('POST', url);
        xhr.onreadystatechange = function() {
            if (xhr.readyState>3 && xhr.status==200) { success(xhr.responseText); }
        };

        xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
        xhr.send(params);
        return xhr;
    }

    function submitForm() {
        let linkElement = document.getElementById("link")
        let nameElement = document.getElementById("name")
        let linkURL = linkElement.value
        let value = nameElement.value

        if (linkURL !== "") {
            post("/api/new", {'originalURL': linkURL, 'value': value}, showLink)
            linkElement.value = ""
            nameElement.value = ""

        } else {
            showMessage()
        }
    }

    function showLink(dataText) {
        let data = JSON.parse(dataText)
        let url = data.GeneratedURL

        let msg = document.getElementById("generated-url");
        let generatedURL = document.getElementById("generated-url-value");

        if (msg.style.display === "none") {
            msg.style.display = "block";
            generatedURL.textContent = url;

        } else {
            generatedURL.textContent = url;
        }
    }

    function showMessage() {
        let messages = document.getElementById("messages-wrap");

        if(messages.classList.contains("hidden")) {
            messages.classList.remove("hidden")
            messages.classList.add("flex")
            document.getElementById("error-message").textContent = emptyInput
        }
    }

    function setAndStoreTheme(styleLink, theme, themeCheckbox) {
        styleLink.setAttribute("href", theme)
        localStorage.setItem(themeCheckbox.id, themeCheckbox.checked)
    }

    function checkThemeSwitchAndUpdateCss(theme) {
        themeSwitch.setAttribute("checked", "checked")
        updateCssLink(theme)
    }

    function uncheckThemeSwitchAndUpdateCss(theme) {
        themeSwitch.removeAttribute("checked")
        updateCssLink(theme)
    }

    function updateCssLink(theme) {
        document.getElementById("style-select").setAttribute("href", theme)
    }
