using Go = import "/go.capnp";
@0xaaa3acc39d94e0f8;
$Go.package("capnpbenchmark");
$Go.import("capnpbenchmark");


struct Account{
    accountId @0 :Data;
    name      @1 :Text;
    balance   @2 :UInt64;
}

struct Result {
    code    @0 : Int32;
    message @1 : Text;
}
struct Response{
    result  @0 : Result;
    account @1 : Account;
}

interface CoreBanking {
  createAccount  @0 (acc :Account) -> (res:Result);
  getAccountInfo @1 (accountId :Data) -> (res :Response);
}

