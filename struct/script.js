function copiarCodigo() {
    const codigo = document.querySelector('.codigo').innerText;
    navigator.clipboard.writeText(codigo).then(function() {
        alert('Código copiado para a área de transferência!');
    }, function(err) {
        console.error('Falha ao copiar o código: ', err);
    });
}
