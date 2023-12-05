package metrics

// Keep the metric fields alphabetized within each category.
const (
	// Common.
	AppVersion       = "app_version"
	AppInfo          = "app_info"
	BlockHeight      = "block_height"
	Count            = "count"
	Detail           = "detail"
	Deterministic    = "deterministic"
	Distribution     = "distribution"
	Error            = "error"
	GitCommit        = "git_commit"
	HttpGet5xx       = "http_get_5xx"
	HttpGetHangup    = "http_get_hangup"
	HttpGetRequest   = "http_get_request"
	HttpGetResponse  = "http_get_response"
	HttpGetTimeout   = "http_get_timeout"
	Invalid          = "invalid"
	Latency          = "latency"
	Matched          = "matched"
	MessageType      = "message_type"
	Msg              = "msg"
	Negative         = "negative"
	NonDeterministic = "non_deterministic"
	Positive         = "positive"
	Reason           = "reason"
	Received         = "received"
	Rejected         = "rejected"
	SampleRate       = "sample_rate"
	SequenceNumber   = "sequence_number"
	Success          = "success"
	Valid            = "valid"
	ValidateBasic    = "validate_basic"
	CheckTx          = "check_tx"
	ReCheckTx        = "recheck_tx"
	DeliverTx        = "deliver_tx"
	ProcessProposal  = "process_proposal"

	// Common (Daemons).
	MainTaskLoop = "main_task_loop"

	// ABCI: Prepare / Process
	AcknowledgeBridgesTx = "acknowledge_bridges_tx"
	ConsensusRound       = "consensus_round"
	DisallowMsg          = "disallow_msg"
	Decode               = "decode"
	FundingTx            = "funding_tx"
	GetTxsInOrder        = "get_txs_in_order"
	Handler              = "handler"
	NumOtherTxs          = "num_other_txs"
	OperationsTx         = "operations_tx"
	OriginalNumTxs       = "original_num_txs"
	OtherTxs             = "other_txs"
	RemoveDisallowMsgs   = "remove_disallow_msgs"
	PrepareProposalTxs   = "prepare_proposal_txs"
	PrepareCheckState    = "prepare_check_state"
	PricesTx             = "prices_tx"
	TotalNumBytes        = "total_num_bytes"
	TotalNumTxs          = "total_num_txs"
	Validate             = "validate"

	// Bridge.
	AcknowledgeBridges            = "acknowledge_bridges"
	AcknowledgedEventInfo         = "acknowledged_event_info"
	BridgeTokenDenom              = "bridge_token_denom"
	CompleteBridge                = "complete_bridge"
	GetAcknowledgeBridges         = "get_acknowledge_bridges"
	LastBridgeEventId             = "last_bridge_event_id"
	LastBridgeEventEthBlockHeight = "last_bridge_event_eth_block_height"
	LastCompletedBridgeId         = "last_completed_bridge_id"
	NextAcknowledgedEventId       = "next_acknowledge_event_id"
	NumBridges                    = "num_bridges"
	UnbridgedBalance              = "unbridged_balance"

	// Bridge Daemon.
	BridgeDaemon = "bridge_daemon"
	NewEthLogs   = "new_eth_logs"

	// Bridge Server.
	AddBridgeEvents          = "add_bridge_events"
	BridgeServer             = "bridge_server"
	EthBlockHeight           = "eth_block_height"
	EventIdAlreadyRecognized = "event_id_already_recognized"
	EventIdNotSequential     = "event_id_not_sequential"
	NextId                   = "next_id"
	RecognizedEventInfo      = "recognized_event_info"

	// CLOB.
	AddPerpetualFillAmount                       = "add_perpetual_fill_amount"
	BaseQuantums                                 = "base_quantums"
	BestAskClobPair                              = "best_ask_clob_pair"
	BestBidClobPair                              = "best_bid_clob_pair"
	Buy                                          = "buy"
	CancelOrder                                  = "cancel_order"
	CancelOrderAccounts                          = "cancel_order_accounts"
	CancelShortTermOrder                         = "cancel_short_term_order"
	CancelStatefulOrder                          = "cancel_stateful_order"
	ClobPairId                                   = "clob_pair_id"
	ClobLiquidateSubaccountsAgainstOrderbook     = "liquidate_subaccounts_against_orderbook"
	LiquidateSubaccounts_GetLiquidations         = "liquidate_subaccounts_against_orderbook_get_liquidations"
	LiquidateSubaccounts_PlaceLiquidations       = "liquidate_subaccounts_against_orderbook_place_liquidations"
	LiquidateSubaccounts_Deleverage              = "liquidate_subaccounts_against_orderbook_deleverage"
	CollateralizationCheck                       = "place_order_collateralization_check"
	CollateralizationCheckFailed                 = "collateralization_check_failed"
	CollateralizationCheckSubaccounts            = "collateralization_check_subaccounts"
	Conditional                                  = "conditional"
	ConditionalOrderTriggered                    = "conditional_order_triggered"
	ConditionalOrderUntriggered                  = "conditional_order_untriggered"
	ConvertToUpdates                             = "convert_to_updates"
	CreateClobPair                               = "create_clob_pair"
	Expired                                      = "expired"
	FullyFilled                                  = "fully_filled"
	GetFillQuoteQuantums                         = "get_fill_quote_quantums"
	Hydrate                                      = "hydrate"
	IsLong                                       = "is_long"
	IterateOverPendingMatches                    = "iterate_over_pending_matches"
	MemClobReplayOperations                      = "memclob_replay_operations"
	MemClobPurgeInvalidState                     = "memclob_purge_invalid_state"
	NumConditionalOrderRemovals                  = "num_conditional_order_removals"
	NumFills                                     = "num_fills"
	NumLongTermOrderRemovals                     = "num_long_term_order_removals"
	NumMatchPerpDeleveragingOperations           = "num_match_perp_deleveraging_operations"
	NumMatchPerpLiquidationsOperations           = "num_match_perp_liquidations_operations"
	NumMatchStatefulOrders                       = "num_match_stateful_orders"
	NumMatchTakerOrders                          = "num_match_taker_orders"
	NumMatchedOrdersInOperationsQueue            = "num_matched_orders_in_operations_queue"
	NumMatchedConditionalOrders                  = "num_match_conditional_orders"
	NumMatchedLiquidationOrders                  = "num_match_liquidation_orders"
	NumMatchedLongTermOrders                     = "num_match_long_term_orders"
	NumMatchedShortTermOrders                    = "num_match_short_term_orders"
	NumOffsettingSubaccountsForDeleveraging      = "num_offsetting_subaccounts_for_deleveraging"
	NumProposedOperations                        = "num_proposed_operations"
	NumShortTermOrderTxBytes                     = "num_short_term_order_tx_bytes"
	NumUniqueSubaccountsDeleveraged              = "num_unique_subaccounts_deleveraged"
	NumUniqueSubaccountsLiquidated               = "num_unique_subaccounts_liquidated"
	NumUniqueSubaccountsOffsettingDeleveraged    = "num_unique_subaccounts_offsetting_deleveraged"
	OffsettingSubaccountPerpetualPosition        = "offsetting_subaccount_perpetual_position"
	OperationsQueueLength                        = "operations_queue_length"
	OrderConflictsWithClobPairStatus             = "order_conflicts_with_clob_pair_status"
	OrderFlag                                    = "order_flag"
	OrderSide                                    = "order_side"
	OrderId                                      = "order_id"
	PartiallyFilled                              = "partially_filled"
	PlaceConditionalOrdersFromLastBlock          = "place_conditional_orders_from_last_block"
	PlaceLongTermOrdersFromLastBlock             = "place_long_term_orders_from_last_block"
	PlaceOrder                                   = "place_order"
	PlaceOrderAccounts                           = "place_order_accounts"
	PlaceStatefulOrder                           = "place_stateful_order"
	ProcessMatches                               = "process_matches"
	ProcessOperations                            = "process_operations"
	ProposedOperations                           = "proposed_operations"
	Proposer                                     = "proposer"
	ProposerConsAddress                          = "proposer_cons_addr"
	QuoteQuantums                                = "quote_quantums"
	RateLimit                                    = "rate_limit"
	ReduceOnly                                   = "reduce_only"
	RemovalReason                                = "removal_reason"
	RemoveAndClearOperationsQueue                = "remove_and_clear_operations_queue"
	ReplayOperations                             = "replay_operations"
	SortLiquidationOrders                        = "sort_liquidation_orders"
	SendCancelOrderOffchainUpdates               = "send_cancel_order_offchain_updates"
	SendPlaceOrderOffchainUpdates                = "send_place_order_offchain_updates"
	SendPlacePerpetualLiquidationOffchainUpdates = "send_perpetual_liquidation_offchain_updates"
	SendPrepareCheckStateOffchainUpdates         = "send_prepare_check_state_offchain_updates"
	SendProcessProposerMatchesOffchainUpdates    = "send_process_proposer_matches_offchain_updates"
	SendProposedOperationsOffchainUpdates        = "send_proposed_operations_offchain_updates"
	SendPurgeOffchainUpdates                     = "send_purge_offchain_updates"
	SendUncrossOffchainUpdates                   = "send_uncross_offchain_updates"
	Sell                                         = "sell"
	ShortTermOrder                               = "short_term_order"
	SkipOrderRemovalAfterPlacement               = "skip_order_removal_after_placement"
	StatefulCancellationMsgHandlerFailure        = "stateful_cancellation_msg_handler_failure"
	StatefulCancellationMsgHandlerSuccess        = "stateful_cancellation_msg_handler_success"
	StatefulOrder                                = "stateful_order"
	StatefulOrderAlreadyRemoved                  = "stateful_order_already_removed"
	StatefulOrderMsgHandlerSuccess               = "stateful_order_msg_handler_success"
	StatefulOrderRemoved                         = "stateful_order_removed"
	Status                                       = "status"
	SubaccountPendingMatches                     = "subaccount_pending_matches"
	TimeInForce                                  = "time_in_force"
	TotalOrdersInClob                            = "total_orders_in_clob"
	TotalQuoteQuantums                           = "total_quote_quantums"
	Unfilled                                     = "unfilled"
	UnfilledLiquidationOrders                    = "unfilled_liquidation_orders"
	UnknownPlaceOrders                           = "unknown_place_orders"
	UnverifiedStatefulOrderRemoval               = "unverified_stateful_order_removal"
	UpdateBlockRateLimitConfiguration            = "update_block_rate_limit_configuration"
	UpdateClobPair                               = "update_clob_pair"
	UpdateEquityTierLimitConfiguration           = "update_equity_tier_limit_configuration"
	UpdateLiquidationsConfig                     = "update_liquidations_config"
	ValidateMatches                              = "validate_matches"
	ValidateOrder                                = "validate_order"

	// MemCLOB.
	AddedToOrderBook                     = "added_to_orderbook"
	AddToOrderbookCollateralizationCheck = "add_to_orderbook_collateralization_check"
	Memclob                              = "memclob"
	RemovedFromOrderBook                 = "removed_from_orderbook"

	// Daemon
	DaemonServer    = "daemon_server"
	ValidResponse   = "valid_response"
	MissingResponse = "missing_response"

	// Epochs.
	EpochInfoName = "epoch_name"
	EpochNumber   = "epoch_number"
	IsEpochOne    = "is_epoch_one"

	// Perpetuals.
	AddPremiumSamples            = "add_premium_samples"
	AddPremiumVotes              = "add_premium_votes"
	GetMarginRequirements        = "get_margin_requirements"
	GetNetNotional               = "get_net_notional"
	GetNotionalInBaseQuantums    = "get_notional_in_base_quantums"
	GetPerpetualAndMarketPrice   = "get_perpetual_and_market_price"
	GetAllPerpetualPricePremiums = "get_all_perpetual_price_premiums"
	NewPremiumVotes              = "new_premium_votes"
	NumPremiumsFromEpoch         = "num_premiums_from_epoch"
	MissingIndexPriceForFunding  = "missing_index_price_for_funding"
	NumPremiumVotes              = "num_premium_votes"
	PerpetualTicker              = "perpetual_ticker"
	PerpetualId                  = "perpetual_id"
	PremiumRate                  = "premium_rate"
	PremiumSampleValue           = "premium_sample_value"
	PremiumType                  = "premium_type"

	// Rewards.
	GetRewardShare                   = "get_reward_share"
	ProcessRewardsForBlock           = "process_rewards_for_block"
	TotalRewardShareWeight           = "total_reward_share_weight"
	DistributedRewardTokens          = "distributed_reward_tokens"
	TreasuryBalanceAfterDistribution = "treasury_balance_after_distribution"

	// Vest.
	GetVestEntry          = "get_vest_entry"
	VestAmount            = "vest_amount"
	BalanceAfterVestEvent = "balance_after_vest_event"
	VesterAccount         = "vester_account"
	ProcessVesting        = "process_vesting"
	AccountTransfer       = "account_transfer"

	// Block Time.
	BlockTimeMs = "block_time_ms"

	// Prices.
	CreateOracleMarket                           = "create_oracle_market"
	CurrentMarketPrices                          = "current_market_prices"
	GetValidMarketPriceUpdates                   = "get_valid_market_price_updates"
	IndexPriceDoesNotExist                       = "index_price_does_not_exist"
	IndexPriceIsZero                             = "index_price_is_zero"
	IndexPriceNotAccurate                        = "index_price_not_accurate"
	IndexPriceNotAvailForAccuracyCheck           = "index_price_not_available_for_accuracy_check"
	LastPriceUpdateForMarketBlock                = "last_price_update_for_market_block"
	MissingPriceUpdates                          = "missing_price_updates"
	NumMarketPricesToUpdate                      = "num_market_prices_to_update"
	PriceChangeRate                              = "price_change_rate"
	ProposedPriceChangesPriceUpdateDecision      = "proposed_price_changes_price_update_decision"
	ProposedPriceCrossesOraclePrice              = "proposed_price_crosses_oracle_price"
	ProposedPriceDoesNotMeetMinPriceChange       = "proposed_price_does_not_meet_min_price_change"
	RecentSmoothedPriceDoesNotMeetMinPriceChange = "recent_smoothed_price_doesnt_meet_min_price_change"
	RecentSmoothedPriceCrossesOraclePrice        = "recent_smoothed_price_crosses_old_price"
	StatefulPriceUpdateValidation                = "stateful_price_update_validation"
	UpdateMarketParam                            = "update_market_param"
	UpdateMarketPrices                           = "update_market_prices"
	UpdateSmoothedPrices                         = "update_smoothed_prices"

	// Sending.
	Account                       = "account"
	New                           = "new"
	ProcessTransfer               = "process_transfer"
	Transfer                      = "transfer"
	ProcessDepositToSubaccount    = "process_deposit_to_subaccount"
	ProcessWithdrawFromSubaccount = "process_withdraw_from_subaccount"
	SendFromModuleToAccount       = "send_from_module_to_account"
	AssetId                       = "asset_id"
	SenderAddress                 = "sender_address"
	SenderModuleName              = "sender_module_name"
	SenderSubaccount              = "sender_subaccount"
	RecipientAddress              = "recipient_address"
	RecipientSubaccount           = "recipient_subaccount"

	// Subaccount.
	CanUpdateSubaccounts                  = "can_update_subaccounts"
	GetNetCollateralAndMarginRequirements = "get_net_collateral_and_margin_requirements"
	GetSubaccount                         = "get_subaccount"
	UpdateSubaccounts                     = "update_subaccounts"
	SubaccountOwner                       = "subaccount_owner"

	// Liquidation Daemon.
	CheckCollateralizationForSubaccounts = "check_collateralization_for_subaccounts"
	GetAllSubaccounts                    = "get_all_subaccounts"
	GetLiquidatableSubaccountIds         = "get_liquidatable_subaccount_ids"
	GetSubaccountsFromKey                = "get_subaccounts_from_key"
	LiquidatableSubaccountIds            = "liquidatable_subaccount_ids"
	LiquidationDaemon                    = "liquidation_daemon"
	PageLimit                            = "page_limit"
	SendLiquidatableSubaccountIds        = "send_liquidatable_subaccount_ids"
	SubaccountsWithOpenPositions         = "subaccounts_with_open_positions"

	// Liquidation.
	ConstructLiquidationOrder             = "construct_liquidation_order"
	InsuranceFundDelta                    = "insurance_fund_delta"
	Liquidations                          = "liquidations"
	MaybeGetLiquidationOrder              = "maybe_get_liquidation_order"
	PlacePerpetualLiquidation             = "place_perpetual_liquidation"
	PercentFilled                         = "percent_filled"
	ProcessLiquidationMatches             = "process_liquidation_matches"
	SubaccountsNotLiquidatable            = "subaccounts_not_liquidatable"
	LiquidationOrderNotionalQuoteQuantums = "liquidation_order_notional_quote_quantums"
	Liquidated                            = "liquidated"
	Filled                                = "filled"
	SubaccountMaxInsuranceLost            = "exceeds_subaccount_max_insurance_lost"
	SubaccountMaxNotionalLiquidated       = "exceeds_subaccount_max_notional_liquidated"
	LiquidationRequiresDeleveraging       = "liquidation_requires_deleveraging"
	LiquidationMatchNegativeTNC           = "liquidation_match_negative_tnc"

	// Deleveraging.
	CannotDeleverageSubaccount     = "cannot_deleverage_subaccount"
	DeleverageSubaccount           = "deleverage_subaccount"
	Deleveraging                   = "deleveraging"
	DeltaQuoteQuantums             = "delta_quote_quantums"
	NumSubaccountsIterated         = "num_subaccounts_iterated"
	NotEnoughPositionToFullyOffset = "not_enough_position_to_fully_offset"
	NonOverlappingBankruptcyPrices = "non_overlapping_bankruptcy_prices"
	NoOpenPositionOnOppositeSide   = "no_open_position_on_opposite_side"

	// Pricefeed Daemon.
	Exchange                                = "exchange"
	ExchangeQueryHandlerApiRequest          = "exchange_query_handler_api_request"
	ExchangeSpecificError                   = "exchange_specific_error"
	GetAllPrices_MarketIdToPrice            = "get_all_prices_market_id_to_price"
	PriceEncoderUpdatePrice                 = "price_encoder_update_price"
	PricefeedDaemon                         = "pricefeed_daemon"
	ConfiguredMarketCount                   = "configured_market_count"
	ConfiguredMarketCountPerExchange        = "configured_market_count_per_exchange"
	ConfiguredExchangeCountPerMarket        = "configured_exchange_count_per_market"
	MarketUpdaterGetAllMarketParams         = "market_updater_get_all_market_params"
	MarketUpdaterApplyMarketUpdates         = "market_updater_apply_market_updates"
	MarketUpdaterUpdateMarkets              = "market_updater_update_markets"
	PriceEncoderPriceConversion             = "price_encoder_price_conversion"
	PriceFetcherQueryExchange               = "price_fetcher_query_exchange"
	PriceFetcherQueryForMarket              = "price_fetcher_query_for_market_sampled"
	PriceFetcherSubtaskLoop                 = "price_fetcher_subtask_loop"
	PriceFetcherSubtaskLoopAndSetCtxTimeout = "price_fetcher_subtask_loop_and_set_ctx_timeout"
	PriceUpdateCount                        = "price_update_count"
	PriceUpdaterSendPrices                  = "price_updater_send_prices"
	PriceUpdaterTaskLoop                    = "price_updater_task_loop"
	PriceUpdaterTransformPrices             = "price_updater_transform_prices"
	PriceUpdaterZeroPrices                  = "price_updater_zero_prices"

	// Pricefeed Server.
	GetValidPrices                = "get_valid_prices"
	ValidPrices                   = "valid_prices"
	NoMarketPrice                 = "no_market_price"
	NoValidMedianPrice            = "no_valid_median_price"
	PricefeedServer               = "pricefeed_server"
	PricefeedServerUpdatePrices   = "pricefeed_server_update_prices"
	PricefeedServerValidatePrices = "pricefeed_server_validate_prices"
	PriceIsInvalid                = "price_is_invalid"

	// Shared Pricefeed Server and Daemon.
	UpdatePrice = "update_price"

	// msgsender
	MessageSendSuccess    = "message_send_success"
	MessageSendError      = "message_send_error"
	SendOffchainData      = "send_offchain_data"
	SendOnchainData       = "send_onchain_data"
	OnchainMessageLength  = "onchain_message_length"
	OffchainMessageLength = "offchain_message_length"

	// Indexer events.
	TotalNumIndexerBlockEvents = "total_num_block_events"
	TotalNumIndexerTxnEvents   = "total_num_txn_events"

	// Mev.
	Mev                            = "mev"
	MevSentDatapoints              = "mev_num_sent_datapoints"
	MidPrice                       = "mid_price"
	MissingMidPrice                = "missing_mid_price"
	ProposerNumFills               = "proposer_num_fills"
	ProposerNumMatchedTakerOrders  = "proposer_num_matched_taker_orders"
	ProposerVolumeQuoteQuantums    = "proposer_volume_quote_quantums"
	ValidatorNumFills              = "validator_num_fills"
	ValidatorNumMatchedTakerOrders = "validator_num_matched_taker_orders"
	ValidatorVolumeQuoteQuantums   = "validator_volume_quote_quantums"
)

const (
	LatencyMetricSampleRate    = 0.01
	AvailableMarketsSampleRate = .1
)
