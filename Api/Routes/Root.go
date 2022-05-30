package Routes

import "net/http"

func Root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>alpha Protocol</title>\n</head>\n<body>\n    <div id=\"reset-ram-wrapper\">\n        <button id=\"reset-ram-btn\" onclick=\"resetRam()\">Reset-RAM</button>\n    </div>\n\n    <p>&nbsp;</p>\n\n    <div id=\"set-time-wrapper\">\n        <label for=\"date-input\">\n            Date and time (2022-04-23T00:09)\n        </label><input id=\"date-input\" type=\"datetime-local\" /> <button id=\"date-btn\" onclick=\"setDateTime()\">set datetime</button>\n    </div>\n\n    <p>&nbsp;</p>\n\n    <div id=\"set-text-wrapper\">\n        <label for=\"position-select\">position</label>\n        <select id=\"position-select\">\n            <option value=\"\">---</option>\n            <option value=\"MIDDLE\">MIDDLE</option>\n            <option value=\"CENTER\">CENTER</option>\n            <option value=\"BOTTOM\">BOTTOM</option>\n            <option value=\"FILL\">FILL</option>\n            <option value=\"LEFT\">LEFT</option>\n            <option value=\"RIGHT\">RIGHT</option>\n        </select>\n\n        <label for=\"default-mode-select\">default-mode</label>\n        <select id=\"default-mode-select\">\n            <option value=\"\">---</option>\n            <option value=\"ROTATE\">ROTATE</option>\n            <option value=\"HOLD\">HOLD</option>\n            <option value=\"FLASH\">FLASH</option>\n            <option value=\"ROLL_UP\">ROLL_UP</option>\n            <option value=\"ROLL_DOWN\">ROLL_DOWN</option>\n            <option value=\"ROLL_LEFT\">ROLL_LEFT</option>\n            <option value=\"ROLL_RIGHT\">ROLL_RIGHT</option>\n            <option value=\"WIPE_UP\">WIPE_UP</option>\n            <option value=\"WIPE_DOWN\">WIPE_DOWN</option>\n            <option value=\"WIPE_LEFT\">WIPE_LEFT</option>\n            <option value=\"WIPE_RIGHT\">WIPE_RIGHT</option>\n            <option value=\"SCROLL\">SCROLL</option>\n            <option value=\"AUTOMODE\">AUTOMODE</option>\n            <option value=\"ROLL_IN\">ROLL_IN</option>\n            <option value=\"ROLL_OUT\">ROLL_OUT</option>\n            <option value=\"WIPE_IN\">WIPE_IN</option>\n            <option value=\"WIPE_OUT\">WIPE_OUT</option>\n            <option value=\"COMPRESSED_ROTATE\">COMPRESSED_ROTATE</option>\n            <option value=\"EXPLODE\">EXPLODE</option>\n            <option value=\"CLOCK\">CLOCK</option>\n            <option value=\"SPECIAL\">SPECIAL</option>\n        </select>\n\n        <label for=\"special-mode-select\">special-mode</label>\n        <select id=\"special-mode-select\">\n            <option value=\"\">---</option>\n            <option value=\"TWINKLE\">TWINKLE</option>\n            <option value=\"SPARKLE\">SPARKLE</option>\n            <option value=\"SNOW\">SNOW</option>\n            <option value=\"INTERLOCK\">INTERLOCK</option>\n            <option value=\"SWITCH\">SWITCH</option>\n            <option value=\"SLIDE\">SLIDE</option>\n            <option value=\"SPRAY\">SPRAY</option>\n            <option value=\"STARBUST\">STARBUST</option>\n            <option value=\"WELCOME\">WELCOME</option>\n            <option value=\"SLOT_MACHINE\">SLOT_MACHINE</option>\n            <option value=\"CYCLE_COLOR\">CYCLE_COLOR</option>\n        </select>\n\n        <label for=\"text\">text</label>\n        <input type=\"text\" id=\"text\">\n\n        <button id=\"text-btn\" onclick=\"setText()\">write text</button>\n    </div>\n\n    <p>&nbsp;</p>\n\n    <div id=\"special-sign-wrapper\">\n        <label for=\"position-select-special\">position</label>\n        <select id=\"position-select-special\">\n            <option value=\"MIDDLE\">MIDDLE</option>\n            <option value=\"CENTER\">CENTER</option>\n            <option value=\"BOTTOM\">BOTTOM</option>\n            <option value=\"FILL\">FILL</option>\n            <option value=\"LEFT\">LEFT</option>\n            <option value=\"RIGHT\">RIGHT</option>\n        </select>\n\n        <label for=\"special-sign-select\">\n            special sign\n        </label>\n        <select id=\"special-sign-select\">\n            <option value=\"THANK_YOU\">THANK_YOU</option>\n            <option value=\"NO_SMOKE\">NO_SMOKE</option>\n            <option value=\"DONT_DRINK_AND_DRIVE\">DONT_DRINK_AND_DRIVE</option>\n            <option value=\"RUNNING_ANIMAIL\">RUNNING_ANIMAIL</option>\n            <option value=\"FIREWORKS\">FIREWORKS</option>\n            <option value=\"TURBO_CAR\">TURBO_CAR</option>\n            <option value=\"CHERRY_BOMB\">CHERRY_BOMB</option>\n        </select>\n\n        <button id=\"special-sign-btn\" onclick=\"setSpecialSign()\">set special sign</button>\n    </div>\n\n    <script type=\"text/javascript\">\n        function resetRam() {\n            httpGetAsync('/api/v1/reset-ram', () => {});\n        }\n\n        function setSpecialSign() {\n            let data = {\n                \"position\": document.getElementById('position-select-special').value,\n                \"special-sign\": document.getElementById('special-sign-select').value\n            }\n\n            httpPostAsync('/api/v1/special-sign', data, () => {})\n        }\n\n        function setDateTime() {\n            let data = {\n                \"date\": document.getElementById('date-input').value\n            }\n\n            httpPostAsync('/api/v1/set-date', data, () => {})\n        }\n\n        function setText() {\n            let data = {\n                \"message\": document.getElementById('text').value,\n                \"position\": document.getElementById('position-select').value === \"\" ? null : document.getElementById('position-select').value,\n                \"default-mode\": document.getElementById('default-mode-select').value === \"\" ? null : document.getElementById('default-mode-select').value,\n                \"special-mode\": document.getElementById('special-mode-select').value === \"\" ? null : document.getElementById('special-mode-select').value,\n            }\n\n            httpPostAsync('/api/v1/write-text', data, () => {})\n        }\n\n        function httpGetAsync(theUrl, callback)\n        {\n            let xmlHttp = new XMLHttpRequest();\n            xmlHttp.onreadystatechange = function() {\n                if (xmlHttp.readyState === 4 && xmlHttp.status === 200)\n                    callback(xmlHttp.responseText);\n            }\n            xmlHttp.open(\"GET\", theUrl, true); // true for asynchronous\n            xmlHttp.send(null);\n        }\n\n        function httpPostAsync(theUrl, theData, callback)\n        {\n            let xhr = new XMLHttpRequest();\n            xhr.open(\"POST\", theUrl, true);\n            xhr.setRequestHeader(\"Content-Type\", \"application/json\");\n            xhr.onreadystatechange = function () {\n                if (xhr.readyState === 4 && xhr.status === 200) {\n                    let json = JSON.parse(xhr.responseText);\n                    callback(json);\n                }\n            };\n            let data = JSON.stringify(theData);\n            xhr.send(data);\n        }\n    </script>\n</body>\n</html>"))
}
