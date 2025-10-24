program genap
kamus
    i, banyak, bilangan, hasil: integer
algoritma
    Input(banyak)
    hasil <-- 0
    for i = 1 to banyak do
        Input(bilangan)
        hasil <-- hasil + (((bilangan mod 2 )- 1)*-1)
    endfor
    Output(hasil)
endprogram