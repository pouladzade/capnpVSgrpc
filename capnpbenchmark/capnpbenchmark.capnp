using Go = import "/go.capnp";
@0xaaa3acc39d94e0f8;
$Go.package("capnpbenchmark");
$Go.import("capnpbenchmark");


struct Account{
    accountId @0 :Data;
    name      @1 :Text;
    balance   @2 :UInt64;
}

interface CoreBanking {
  echoAccount    @0 (acc1 :Account) -> (acc2 :Account);
}

