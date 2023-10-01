package limecosdk

type AccountDataAction string

const SubscribeBalance AccountDataAction = "subscribeBalance"
const SubscribePositions AccountDataAction = "subscribeBalance"
const SubscribeOrders AccountDataAction = "subscribeBalance"
const SubscribeTrades AccountDataAction = "subscribeBalance"

const UnsubscribeBalance AccountDataAction = "subscribeBalance"
const UnsubscribePositions AccountDataAction = "subscribeBalance"
const UnsubscribeOrders AccountDataAction = "subscribeBalance"
const UnsubscribeTrades AccountDataAction = "subscribeBalance"

type MarketDataAction string

const Subscribe MarketDataAction = "subscribe"
const Unsubscribe MarketDataAction = "unsubscribe"
