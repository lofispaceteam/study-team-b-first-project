document.addEventListener('DOMContentLoaded', () =>{
    loadAllMemes();

    const bSwitch = document.querySelector('.switch');
    if (bSwitch){
        bSwitch.addEventListener('click', toggleTheme);

        const savedTheme = localStorage.getItem('theme');
        if (savedTheme == 'dark'){
            document.body.classList.add('dark');
            bSwitch.textContent = '☾';
        }
    }
});

async function loadAllMemes(){
    const container = document.querySelector('.container');
    if (container){
        try{
            const responce = await fetch('http://localhost:8081/api/memes');
            if (!responce.ok) throw new Error('Failed to fetch memes');
            const memes = await responce.json();
            
            container.innerHTML = '';
            memes.forEach(meme => {
                const memeDiv = document.createElement('div');
                memeDiv.className = 'meme-item';
                memeDiv.innerHTML = `
                    <img src="${meme.url}" alt="${meme.text}">
                    <p>${meme.text}</p>
                `;
                container.appendChild(memeDiv);
            });
        } catch (error){
            console.error(error);
            container.innerHTML = `<p>Ошбика загрузки мемов</p>`;
        }
    }
    else{
        alert(" d  ");
    }
}

function toggleTheme(){
    const bSwitch = document.querySelector('.switch');
    const body = document.body.classList;
    body.toggle('dark');
    bSwitch.textContent = body.contains('dark') ? '☾' : '☀︎';

    localStorage.setItem('theme', body.contains('dark' ? 'dark' : 'light'));
}