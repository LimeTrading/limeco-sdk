package limecosdk

type AccountDataAction string

// Subscribe to an accounts balance
const SubscribeBalance AccountDataAction = "subscribeBalance"

// Subscribe to account positions
const SubscribePositions AccountDataAction = "subscribeBalance"

// Subscribe to account orders
const SubscribeOrders AccountDataAction = "subscribeBalance"

// Subscribe to account transactions
const SubscribeTrades AccountDataAction = "subscribeBalance"

// Unsubscribe to an accounts balance
const UnsubscribeBalance AccountDataAction = "subscribeBalance"

// Unsubscribe to account positions
const UnsubscribePositions AccountDataAction = "subscribeBalance"

// Unsubscribe to account orders
const UnsubscribeOrders AccountDataAction = "subscribeBalance"

// Unsubscribe to account transactions
const UnsubscribeTrades AccountDataAction = "subscribeBalance"

type MarketDataAction string

// Subscribe to market data
const Subscribe MarketDataAction = "subscribe"

// Unsubscribe from market data
const Unsubscribe MarketDataAction = "unsubscribe"
