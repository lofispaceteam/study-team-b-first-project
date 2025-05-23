document.addEventListener('DOMContentLoaded', () =>{
    loadAllMemes();

    const bSwitch = document.querySelector('.switch');
    if (bSwitch){
        bSwitch.addEventListener('click', toggleTheme);

        const savedTheme = localStorage.getItem('theme');
        if (savedTheme == 'dark'){
            document.body.classList.add('dark');
            bSwitch.textContent = '☀︎';
        }
    }

    document.getElementById('btn-close').addEventListener('click', closeModal)

    document.getElementById('card-meme').addEventListener('click', (e) => {
        if (e.target === document.getElementById('card-meme')) {
            closeModal();
        }
    });

    document.getElementById('btn-addmeme').addEventListener('click', () => openModal(null, null, null));

    document.getElementById('btn-copy').addEventListener('click', async () => {
        try{
            const text = document.getElementById('input-meme-path').value;
            await navigator.clipboard.writeText(text);
        } catch(error){
            console.log(error);
            alert('Не удалось скопировать в буфер обмен');
        }
    });

    document.getElementById('btn-paste').addEventListener('click', async () => {
        try{
            const text = await navigator.clipboard.readText();
            document.getElementById('input-meme-path').value = text;
            document.getElementById('img-modal').src = text;
        } catch(error){
            console.log(error);
            alert('Не удалось вставить данные из буфера обмена');
        };
    });

    const path = document.getElementById('input-meme-path');
    const caption = document.getElementById('input-meme-text');
    document.getElementById('btn-modal-addmeme').addEventListener('click', async () => {
        const url = path.value;
        const text = caption.value;
        if (!url || !text){
            alert('Поля должны быть заполнеными!');
            return;
        }
        try{
            const responce = await fetch('http://localhost:8080/api/memes', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json'},
                body: JSON.stringify({"image_url": url, "title": text})
            })
            alert("Успешно вставили вашу картинку, добавите ещё?");
            path.value = '';
            caption.value = '';
            closeModal();
            loadAllMemes();
        } catch(error){
            alert('Ошибка добавления мема');
            console.error(error);
        }

    });

    document.getElementById('btn-edit').addEventListener('click', async () => {
        const url = path.value;
        const text = caption.value;
        const memeId = path.dataset.memeId;
        if (!url || !text || !memeId){
            alert('Поля должны быть заполнеными!');
            return;
        }
        try{
            const responce = await fetch(`http://localhost:8080/api/memes/${memeId}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json'},
                body: JSON.stringify({'image_url': url, 'title': text})
            })
            alert('Изменения вступили в силу!');
            closeModal();
            loadAllMemes();
        } catch(error){
            alert("Ошибка изменения мема!");
            console.error(error);
        }
    });

    document.getElementById('btn-delete').addEventListener('click', async () => {
        const memeId = path.dataset.memeId;
        try{
            const responce = await fetch(`http://localhost:8080/api/memes/${memeId}`, {
                method: 'DELETE',
                headers: { 'Content-Type': 'application/json'}
            });
            alert('Мем успешно удалён!');
            closeModal();
            loadAllMemes();
        } catch(error){
            alert('Ошибка удаления мема!');
            console.error(error);
        }
    });

    
});

async function loadAllMemes(){
    const container = document.querySelector('.container');
    if (container){
        try{
            const responce = await fetch('http://localhost:8080/api/memes');
            if (!responce.ok) throw new Error('Failed to fetch memes');
            const memes = await responce.json();
            
            container.innerHTML = '';
            memes.forEach(meme => {
                const memeDiv = document.createElement('div');
                memeDiv.className = 'meme-item';
                memeDiv.innerHTML = `
                    <img src="${meme.image_url}" alt="${meme.title}">
                    <p>${meme.title}</p>
                `;
                container.appendChild(memeDiv);
                memeDiv.addEventListener('click', () => openModal(meme.id, meme.image_url, meme.title));
            });
        } catch (error){
            console.error(error);
            container.innerHTML = `<p>Ошибка загрузки мемов</p>`;
        }
    }
}

function toggleTheme(){
    const bSwitch = document.querySelector('.switch');
    const body = document.body.classList;
    body.toggle('dark');
    bSwitch.textContent = body.contains('dark') ? '☀︎' : '☾';

    localStorage.setItem('theme', body.contains('dark' ? 'dark' : 'light'));
}

function openModal(id, url, text){
    const modal = document.getElementById('card-meme');
    const image = document.getElementById('img-modal');
    const path = document.getElementById('input-meme-path');
    const caption = document.getElementById('input-meme-text');
    const btns = document.getElementById('buttons');
    const btnAdd = document.getElementById('btn-modal-addmeme');

    if (url == null || text == null){
        btns.style.display = 'none';
        btnAdd.style.display = 'flex';
        path.dataset.memeId = '';
    }
    else{
        btns.style.display = 'flex';
        btnAdd.style.display = 'none';
        path.dataset.memeId = id;
    }
    image.src = url;
    path.value = url;
    caption.value = text;
    modal.style.display = 'flex';
}

function closeModal(){
    document.getElementById('card-meme').style.display = 'none';
}