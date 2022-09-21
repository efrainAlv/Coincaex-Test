var account;

function connect(event) {
    let button = event.target;
    ethereum.request({ method: 'eth_requestAccounts' }).then(accounts => {
        account = accounts[0];
        console.log(account);
        button.textContent = account;

        ethereum.request({ method: 'eth_getBalance', params: [account, 'latest'] }).then(result => {
            console.log(result);
            let wei = parseInt(result, 16);
            let balance = wei / (10 ** 18);
            console.log(balance + " ETH");
        });
    });
}


function send() {
    let transactionParam = {
        to: '0x45B6b39e1Cf8A6b4Ff2720f6BA0089d4574126E5',
        from: account,
        value: '0x38D7EA4C68000'
    };

    ethereum.request({ method: 'eth_sendTransaction', params: [transactionParam] }).then(txhash => {
        console.log(txhash);
        checkTransactionconfirmation(txhash).then(r => alert(r));
    });
}

function checkTransactionconfirmation(txhash) {

    let checkTransactionLoop = () => {
        return ethereum.request({ method: 'eth_getTransactionReceipt', params: [txhash] }).then(r => {
            if (r != null) return 'confirmed';
            else return checkTransactionLoop();
        });
    };

    return checkTransactionLoop();
}