import csv
from pathlib import Path

PATH = Path('data/UhttBarcodeReference-master/DATA')
OUTPUT_SQL = Path('data/UhttBarcodeReference-master/out.sql')


def main():
    with OUTPUT_SQL.open('w') as fout:
        for path in PATH.glob('*.csv'):
            if not path.stem.split('_')[-1].isnumeric():
                print('ignored', path)
                continue

            print(path)
            with path.open() as f:
                reader = csv.reader(f, delimiter='\t')
                next(reader, None)  # skip the headers
                for row in reader:
                    # print(row)
                    fout.write(
                        "INSERT INTO monica_storage.public.barcode"
                        " (type_id, code, name, source_id)"
                        "VALUES (1, '{}', '{}', 1);\n".format(
                            escape(row[1]), escape(row[2])
                        )
                    )


def escape(val):
    return val.replace("'", "''")


if __name__ == '__main__':
    main()
