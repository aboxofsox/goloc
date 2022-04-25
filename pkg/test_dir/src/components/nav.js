const listItems = [
    {link: '#', text: 'Home'},
    {link: '#', text: 'About'},
    {link: '#', text: 'Contact'},
    {link: '#', text: 'Login'},
]

const nav = document.createElement('nav')

listItems.forEach(item => {
    const li = document.createElement('li')
    const a = document.createElement('a')
    a.href = item.link
    a.textContent = item.text
    li.appendChild(a)
    nav.appendChild(li)
})