# Простой генератор паролей написанный на Go
Минимальная длина пароля: 12 символов

# Аргументы программы
Возможность сгенерировать сразу несколько паролей, кол-во должно быть положительным числом.
```bash
./password-generator

./password-generator -help

./password-generator -c 5

./password-generator -count 10
```

# Полезная информация
## Работа с Makefile в NixOS
```bash
nix-shell -p gnumake
```
