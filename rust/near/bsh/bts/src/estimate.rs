use near_sdk::{Balance, Gas};

pub const GAS_FOR_RESOLVE_TRANSFER: Gas = Gas(10_000_000_000_000);
pub const GAS_FOR_FT_TRANSFER_CALL: Gas = Gas(25_000_000_000_000);
pub const GAS_FOR_MINT: Gas = Gas(10_000_000_000_000);
pub const GAS_FOR_ON_MINT: Gas = Gas(10_000_000_000_000);
pub const GAS_FOR_SEND_SERVICE_MESSAGE: Gas = Gas(15_000_000_000_000);
pub const NO_DEPOSIT: Balance = 0;
pub const GAS_FOR_BURN: Gas = Gas(1_000_000_000_000);
pub const GAS_FOR_TOKEN_STORAGE_DEPOSIT: Gas = Gas(8_000_000_000_000);
pub const ONE_YOCTO: Balance = 1;
