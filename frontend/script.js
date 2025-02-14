function consultarCEP() {
    const cep = document.getElementById('cep').value;
    const resultadoDiv = document.getElementById('resultado');

    if (cep.length !== 8 || isNaN(cep)) {
        resultadoDiv.innerHTML = '<p style="color: red;">CEP inválido. Digite 8 números.</p>';
        return;
    }

    fetch(`http://localhost:8080/cep/${cep}`)
        .then(response => {
            if (!response.ok) {
                throw new Error('CEP não encontrado');
            }
            return response.json();
        })
        .then(data => {

            resultadoDiv.innerHTML = `
                <p><strong>CEP:</strong> ${data.cep}</p>
                <p><strong>Logradouro:</strong> ${data.logradouro}</p>
                <p><strong>Bairro:</strong> ${data.bairro}</p>
                <p><strong>Cidade:</strong> ${data.localidade}</p>
                <p><strong>Estado:</strong> ${data.uf}</p>
            `;
        })
        .catch(error => {
            resultadoDiv.innerHTML = `<p style="color: red;">${error.message}</p>`;
        });
}