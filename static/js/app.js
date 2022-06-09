let BASE_URL = `/`;
const Api = {
        shortenUrl: async n => {
            const e = {
                    method: "GET"
                };
                longUrl = encodeURI(n.url)
                t = await fetch(  BASE_URL + `api/v1/encode?longUrl=${longUrl}`, e),
                o = await t.json();
            if (t.ok) return o;
            throw new Error(`${o.message}`)
        }
    },
    form = document.querySelector("form");
let urlList = document.getElementById("url-list");

function addUrlEntry(n) {
    const e = DOMPurify.sanitize(n),
        t = document.createElement("div");
    t.classList.add("py-2", "flex", "justify-between", "text-su-fg-1", "dark:text-su-dark-fg-1");
    const o = `
    <a href="/${e}">
      ${window.location.host}/${e}
    </a>

    <div class="flex space-x-2.5 text-sm">
      <button type="button">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
        </svg>
      </button>
    </div>
  `;
    t.innerHTML = o, t.addEventListener("click", () => {
      navigator.clipboard.writeText(`https://${window.location.host}/${e}`);
      setTimeout(()=> alert("Copied to clipboard"));
    }), urlList.prepend(t)
}
form?.addEventListener("submit", n => {
    n.preventDefault();
    let e = {};
    new FormData(document.querySelector("form")).forEach((t, o) => e[o] = t), Api.shortenUrl(e).then(t => {
      if (n.message) {
        return alert(n.message);
      }
        urlList.classList.contains("hidden") && urlList.classList.remove("hidden");
        let o = document.getElementById("url");
        form.reset(), o.focus(), 
        addUrlEntry(t.tiny_url)
    }).catch(t => {
        alert(`I ran into a problem!
${t}`)
    })
});