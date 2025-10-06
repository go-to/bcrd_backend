drop function if exists sf_translate_case(text);
create function sf_translate_case(
    p_src text
) returns text as
$$
declare
    text_result                  text   := upper(p_src);
    arr_hankaku_daku_kanas       text[] := array ['ｳﾞ', 'ｶﾞ', 'ｷﾞ', 'ｸﾞ', 'ｹﾞ', 'ｺﾞ', 'ｻﾞ', 'ｼﾞ', 'ｽﾞ', 'ｾﾞ', 'ｿﾞ', 'ﾀﾞ', 'ﾁﾞ', 'ﾂﾞ', 'ﾃﾞ', 'ﾄﾞ', 'ﾊﾞ', 'ﾋﾞ', 'ﾌﾞ', 'ﾍﾞ', 'ﾎﾞ', 'ﾊﾟ', 'ﾋﾟ', 'ﾌﾟ', 'ﾍﾟ', 'ﾎﾟ'];
    arr_zenkaku_daku_kanas       text[] := array ['ヴ', 'ガ', 'ギ', 'グ', 'ゲ', 'ゴ', 'ザ', 'ジ', 'ズ', 'ゼ', 'ゾ', 'ダ', 'ヂ', 'ヅ', 'デ', 'ド', 'バ', 'ビ', 'ブ', 'ベ', 'ボ', 'パ', 'ピ', 'プ', 'ペ', 'ポ'];
    text_hankaku_kanas           text   := 'ｱｲｳｴｵｶｷｸｹｺｻｼｽｾｿﾀﾁﾂﾃﾄﾅﾆﾇﾈﾉﾊﾋﾌﾍﾎﾏﾐﾑﾒﾓﾔﾕﾖﾗﾘﾙﾚﾛﾜｦﾝｧｨｩｪｫｯｬｭｮ';
    text_zenkaku_hiras           text   := 'あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをんぁぃぅぇぉっゃゅょ';
    text_zenkaku_kanas           text   := 'アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲンァィゥェォッャュョ';
    text_zenkaku_daku_hiras      text   := 'がぎぐげござじずぜぞだぢづでどばびぶべぼぱぴぷぺぽ';
    text_zenkaku_daku_kanas      text   := 'ガギグゲゴザジズゼゾダヂヅデドバビブベボパピプペポ';
    text_hankaku_symbols         text   := '｡｢｣､･ｰ－';
    text_zenkaku_symbols         text   := '。「」、・ー-';
    text_zenkaku_space           text   := '　';
    text_hankaku_space           text   := ' ';
    text_zenkaku_nums            text   := '０１２３４５６７８９';
    text_hankaku_nums            text   := '0123456789';
    text_zenkaku_upper_alphabets text   := 'ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ';
    text_zenkaku_lower_alphabets text   := 'ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ';
    text_hankaku_upper_alphabets text   := 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
begin
    if p_src is null or p_src = '' then
        return p_src;
    end if;

    for i in 1..array_length(arr_hankaku_daku_kanas, 1)
        loop
            text_result := replace(text_result, arr_hankaku_daku_kanas[i], arr_zenkaku_daku_kanas[i]);
        end loop;

    text_result := translate(
            text_result
        , text_hankaku_kanas || text_zenkaku_hiras || text_zenkaku_daku_hiras || text_hankaku_symbols ||
          text_zenkaku_space || text_zenkaku_nums || text_zenkaku_upper_alphabets || text_zenkaku_lower_alphabets
        , text_zenkaku_kanas || text_zenkaku_kanas || text_zenkaku_daku_kanas || text_zenkaku_symbols ||
          text_hankaku_space || text_hankaku_nums || text_hankaku_upper_alphabets || text_hankaku_upper_alphabets
                   );

    return text_result;
end;
$$ language plpgsql immutable
