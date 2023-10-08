# Simple_Go_Blockchain
Blockchain-Implementierung in Go
Dies ist eine einfache Implementierung einer Blockchain in Go. 
Die Blockchain besteht aus Blöcken, die jeweils eine bestimmte Anzahl von Transaktionen enthalten. 
Jeder Block wird durch einen kryptografischen Hash identifiziert, der aus dem Hash des vorherigen Blocks,
den Transaktionsdaten und einem Proof-of-Work (PoW)-Wert berechnet wird.

Voraussetzungen
Go v1.16 oder höher
Installation
Klonen Sie dieses Repository mit git clone https://github.com/USERNAME/blockchain-go.git.
Navigieren Sie in den Projektordner cd blockchain-go.
Führen Sie den Befehl go run main.go aus, um die Beispielausgabe zu sehen.
Verwendung
Die Implementierung besteht aus zwei benutzerdefinierten Typen: Block und Blockchain.

Block
Der Block-Typ enthält folgende Felder:

data : Ein map[string]interface{} mit den Transaktionsdetails (Absender, Empfänger und Überweisungsbetrag).
hash : Ein String, der den Hash des aktuellen Blocks enthält.
previousHash : Ein String, der den Hash des vorherigen Blocks enthält.
timestamp : Ein time.Time -Objekt, das den Zeitstempel des Blocks enthält.
pow : Eine Ganzzahl, die den aktuellen PoW-Wert des Blocks enthält.
Der Block-Typ verfügt über folgende Methoden:

CalculateHash() string : Diese Methode berechnet den Hash-Wert des Blocks aus den obigen Feldern und gibt ihn als Hex-String zurück.
Mine(difficulty int) : Diese Methode erhöht den PoW-Wert des Blocks, bis ein gültiger Hash gefunden wird (d.h. ein Hash, der mit einer bestimmten Anzahl von Nullen beginnt).
Blockchain
Der Blockchain-Typ enthält folgende Felder:

genesisBlock : Ein Block-Objekt, das den ersten Block der Blockchain darstellt (auch bekannt als Genesis-Block).
chain : Eine Liste von Block-Objekten, die die gesamte Blockchain darstellen.
difficulty : Eine Ganzzahl, die angibt, wie viele Nullen der Hash eines Blocks enthalten muss, um als gültiger Block akzeptiert zu werden.
Der Blockchain-Typ verfügt über folgende Methoden:

AddBlock(from string, to string, amount float64) : Diese Methode fügt einen neuen
