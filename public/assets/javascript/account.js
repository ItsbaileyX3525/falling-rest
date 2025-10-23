const username = document.getElementById("username");
import { returnAPIResponse } from "./api.js";


async function fetchEmail() {
    const response = await fetch("/auth/me", {
        method: "GET",
        credentials: "include",
        headers: {
            "Accept" : "application/json"
        }
    })

    if (!response.ok) {
        console.log("Something went wrong fetching email");
        return
    }

    const data = await response.json()

    console.log(data)
    if (data.success || data.success === "true") {
        return data.email
    } else {
        return "Not logged in"
    }
}

document.addEventListener("DOMContentLoaded", async () => {
    username.innerText = await fetchEmail()

    await returnAPIResponse("/api/decode?input=dGVzdA==?type=base64?apiKey=abc123");
})