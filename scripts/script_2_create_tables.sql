CREATE TABLE IF NOT EXISTS public.barcode_type
(
    id   INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name varchar
);

CREATE TABLE IF NOT EXISTS public.barcode_source
(
    id   INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name varchar
);

CREATE TABLE IF NOT EXISTS public.barcode
(
    id     INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    type_id   INTEGER,
    code   VARCHAR,
    name   VARCHAR,
    source_id INTEGER,

    CONSTRAINT fk_barcode_type
        FOREIGN KEY (type_id)
            REFERENCES barcode_type (id),

    CONSTRAINT fk_barcode_source
        FOREIGN KEY (source_id)
            REFERENCES barcode_source (id)
);

INSERT INTO public.barcode_type (name) values ('UPC/EAN');  -- id=1
INSERT INTO public.barcode_source (name) values ('UhttBarcodeReference');  -- id=1
