const form = document.getElementById("form-tp");
const lista = document.getElementById("lista-tps");

function cargarTPs() {
    fetch("/tps")
        .then(response => response.json())
        .then(tps => {
            lista.innerHTML = "";
            if (!tps) return;

            tps.forEach(tp => {
                const li = document.createElement("li");
                li.classList.add("item-tp");

                const fecha = tp.fecha_entrega.split("T")[0];
                li.textContent = tp.titulo + " - " + tp.materia + " (" + fecha + ") [" + tp.estado + "] ";

                const btn = document.createElement("button");
                btn.textContent = "Eliminar";
                btn.classList.add("btn-eliminar");
                btn.setAttribute("data-id", tp.id);

                li.appendChild(btn);
                lista.appendChild(li);
            });
        });
}

form.addEventListener("submit", (e) => {
    e.preventDefault();

    const nuevoTP = {
        titulo: document.getElementById("titulo").value,
        materia: document.getElementById("materia").value,
        fecha_entrega: document.getElementById("fecha_entrega").value,
        estado: document.getElementById("estado").value
    };

    fetch("/tps", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(nuevoTP)
    }).then(() => {
        form.reset();
        cargarTPs();
    });
});

lista.addEventListener("click", (e) => {
    if (e.target.classList.contains("btn-eliminar")) {
        const id = e.target.getAttribute("data-id");
        fetch("/tps/" + id, {
            method: "DELETE"
        }).then(() => {
            cargarTPs();
        });
    }
});

document.addEventListener("DOMContentLoaded", cargarTPs);
