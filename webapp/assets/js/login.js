$('#login').on('submit', fazerLogin);

function fazerLogin(evento) {
    evento.preventDefault();

    console.log('aqui');

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#email').val(),
            senha: $('#senha').val()
        }
    }).done(function() {
        console.log('kevem')
        window.location = "/home";
    }).fail(function() {
        console.log('error')
        Swal.fire('Ops...', 'Usuário ou senha inválidos!', 'error');
    });
}