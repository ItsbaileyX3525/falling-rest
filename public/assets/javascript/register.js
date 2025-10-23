const form = document.getElementById("form")

/*
curl -X POST http://localhost:8081/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"pass123"}'
*/

form.addEventListener("submit", async (e) => {
    e.preventDefault()
    const formData = new FormData(form)
    
    const response = await fetch("/auth/register", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        credentials: "include",
        body: JSON.stringify({ 
            "email": formData.get("email"),
            "password": formData.get("pwd")
        })
    })

    if (!response.ok) {
        console.log("something went wrong")
        return
    }

    const data = await response.json()

    if (data.success || data.success === 'true') {
        window.location.href = "/account"
    } else {
        console.log(data.message)
    }
})

document.addEventListener("DOMContentLoaded", async () => {
    
})