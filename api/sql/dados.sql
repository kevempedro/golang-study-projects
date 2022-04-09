INSERT INTO usuarios (nome, nick, email, senha)
VALUES
("Kevem Pedro", "kevempedro", "kevempedro@gmail.com", "$2a$10$pWSD4Yq3RnhtoBIwaGUUmeEXtpapo4E1D2AX6S2KEh4CA92WPOPuK"),
("Kevem Mesquita", "kevemmesquita", "kevemmesquita@gmail.com", "$2a$10$pWSD4Yq3RnhtoBIwaGUUmeEXtpapo4E1D2AX6S2KEh4CA92WPOPuK"),
("Kevem Lima", "kevemlima", "kevemlima@gmail.com", "$2a$10$pWSD4Yq3RnhtoBIwaGUUmeEXtpapo4E1D2AX6S2KEh4CA92WPOPuK");

INSERT INTO seguidores (usuario_id, seguidor_id)
VALUES
(1, 2),
(3, 1),
(1, 3);

INSERT INTO publicacoes(titulo, conteudo, autor_id)
VALUES
("Publicação do usuário 1", "Essa é uma publicação do usuário 1", 1),
("Publicação do usuário 2", "Essa é uma publicação do usuário 2", 2),
("Publicação do usuário 3", "Essa é uma publicação do usuário 3", 3);