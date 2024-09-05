class User {
    constructor(name, age, email) {
        this.name = name
        this.age = age
        this.email = email
    }
}

document.onload = () => {
    console.log("LOAD")
    fillTable();
}

const root = document.querySelector("#root")
const usersTable = document.querySelector("#users__table")
let users = []

const fillTable = () => {
    usersTable.innerHTML = `
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Age</th>
                    <th>Email</th>
                </tr>
            </thead>`
}

const loadUsers = () => {
    let jsonData = []
    fetch("http://localhost:8080/users")
        .then(res => res.json())
        .then(userData => users = userData.map(user => new User(user.name, user.age, user.email)))
        .then(() => render(users))
        .catch(err => console.log(err))
}

const loadUser = (email) => {
    fetch(`http://localhost:8080/users/?email=${email}`)
        .then(res => res.json())
        .then(userData => users = [new User(userData.name, userData.age, userData.email)])
        .then(() => render(users))
        .catch(err => console.log(err))

    if(users.length > 0) {
        render(users)
    }
}

const deleteUser = (email) => {
    fetch(`http://localhost:8080/users/delete/?email=${email}`, {
        method: "DELETE"
    })
        .then(res => {if(res.ok) {
            alert("Deleted")
        }})
        .catch(err => console.log(err))
}


const render = (users) => {

    fillTable();

    users.forEach(user => {
        const row = document.createElement("tr")

        row.innerHTML = `
            <td>${user.name}</td>
            <td>${user.age}</td>
            <td>${user.email}</td>`

        usersTable.appendChild(row)
    })
}

const loadUserFrom = document.querySelector(".form")
loadUserFrom.addEventListener("submit", (e) => {
    e.preventDefault()
})

const loadUsersBtn = document.querySelector(".btn-load__users")
loadUsersBtn.addEventListener("click", loadUsers)

const loadUserByEmailBtn = document.querySelector(".btn-load__user")
loadUserByEmailBtn.addEventListener("click", (e) => {
    loadUser(e.target.closest("form").email.value);
})

const deleteUserBtn = document.querySelector(".btn-delete__user")
deleteUserBtn.addEventListener("click", (e) => {
    deleteUser(e.target.closest("form").email.value);
})


