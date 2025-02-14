

# **Consulta de CEP com ViaCEP**

Um projeto simples para consultar informações de endereço a partir de um CEP, utilizando um **backend em Go** e um **frontend em HTML/JavaScript**.

---

## **Funcionalidades**

- Consulta de CEP utilizando a API do [ViaCEP](https://viacep.com.br/).
- Exibição dos dados do endereço (logradouro, bairro, cidade, estado, etc.).
- Interface simples e intuitiva no frontend.

---

## **Tecnologias Utilizadas**

### **Backend**
- **Golang** (Go): Linguagem de programação.
- **Gin**: Framework web para criar a API.
- **ViaCEP**: API pública para consulta de CEPs.

### **Frontend**
- **HTML**: Estrutura da página.
- **CSS**: Estilização básica.
- **JavaScript**: Lógica para fazer requisições à API e exibir os dados.

---

## **Como Executar o Projeto**

### **Pré-requisitos**
- Go instalado (versão 1.16 ou superior).
- Navegador moderno (Chrome, Firefox, Edge, etc.).

### **Passos para Executar**

1. **Clone o repositório**:
   ```bash
   git clone https://github.com/seu-usuario/viacep-com-frontend.git
   cd viacep-com-frontend
   ```

2. **Execute o Backend**:
   - Navegue até a pasta `backend`:
     ```bash
     cd backend
     ```
   - Instale as dependências:
     ```bash
     go mod tidy
     ```
   - Inicie o servidor:
     ```bash
     go run main.go
     ```
   - O backend estará rodando em `http://localhost:8080`.

3. **Abra o Frontend**:
   - Navegue até a pasta `frontend`:
     ```bash
     cd ../frontend
     ```
   - Abra o arquivo `index.html` no navegador.

4. **Consulte um CEP**:
   - Insira um CEP válido (apenas números) no campo de texto e clique em "Consultar".
   - Os dados do endereço serão exibidos abaixo.

---

## **Estrutura do Projeto**

```
viacep-com-frontend/
├── backend/
│   ├── main.go           # Código do backend em Go
│   ├── go.mod            # Dependências do Go
│   └── go.sum            # Dependências do Go
├── frontend/
│   ├── index.html        # Página HTML do frontend
│   └── script.js         # Lógica JavaScript do frontend
└── README.md             # Documentação do projeto
```

---

## **Exemplo de Uso**

1. Abra o frontend no navegador.
2. Insira o CEP `01001000` e clique em "Consultar".
3. O resultado será:

   ```
   CEP: 01001-000
   Logradouro: Praça da Sé
   Bairro: Sé
   Cidade: São Paulo
   Estado: SP
   ```

---
![image](https://github.com/user-attachments/assets/6442a769-af6a-4f53-a1de-fa956d804514)

## **Contribuição**

Contribuições são bem-vindas! Siga os passos abaixo:

1. Faça um fork do projeto.
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`).
3. Commit suas mudanças (`git commit -m 'Adiciona nova feature'`).
4. Push para a branch (`git push origin feature/nova-feature`).
5. Abra um Pull Request.

---

## **Licença**

Este projeto está licenciado sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---

## **Autor**

Feito com ❤️ por [Gilberto Jr](https://github.com/infinitytec15).

---

### **Links Úteis**
- [ViaCEP](https://viacep.com.br/)
- [Documentação do Go](https://golang.org/doc/)
- [Documentação do Gin](https://gin-gonic.com/docs/)

