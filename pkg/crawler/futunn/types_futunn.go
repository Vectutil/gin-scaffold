package futunn

type Futunn struct {
	StockInfo struct {
		Name                         string `json:"name"`
		EnName                       string `json:"enName"`
		MarketType                   int    `json:"marketType"`
		MarketLabel                  string `json:"marketLabel"`
		IsPlate                      bool   `json:"isPlate"`
		StockCode                    string `json:"stockCode"`
		StockId                      string `json:"stockId"`
		MarketCode                   int    `json:"marketCode"`
		InstrumentType               int    `json:"instrumentType"`
		LotSize                      int    `json:"lotSize"`
		SpreadCode                   string `json:"spreadCode"`
		PriceAccuracy                int    `json:"priceAccuracy"`
		SubInstrumentType            int    `json:"subInstrumentType"`
		IsOption                     bool   `json:"isOption"`
		IsFutures                    bool   `json:"isFutures"`
		PriceNominal                 string `json:"priceNominal"`
		PriceLastClose               string `json:"priceLastClose"`
		ServerSendToClientTimeMs     string `json:"serverSendToClientTimeMs"`
		ExchangeDataTimeMs           string `json:"exchangeDataTimeMs"`
		ServerRecvFromExchangeTimeMs string `json:"serverRecvFromExchangeTimeMs"`
		PriceOpen                    string `json:"priceOpen"`
		PriceHighest                 string `json:"priceHighest"`
		PriceLowest                  string `json:"priceLowest"`
		Volume                       string `json:"volume"`
		VolumeSold                   string `json:"volumeSold"`
		VolumeBought                 string `json:"volumeBought"`
		Turnover                     string `json:"turnover"`
		RatioVolume                  string `json:"ratioVolume"`
		RatioTurnover                string `json:"ratioTurnover"`
		AmplitudePrice               string `json:"amplitudePrice"`
		PriceAverage                 string `json:"priceAverage"`
		ChangeSpeedPrice             string `json:"changeSpeedPrice"`
		RatioBidAsk                  string `json:"ratioBidAsk"`
		VolumePrecision              int    `json:"volumePrecision"`
		TotalShares                  string `json:"totalShares"`
		TotalMarketCap               string `json:"totalMarketCap"`
		OutstandingShares            string `json:"outstandingShares"`
		OutstandingMarketCap         string `json:"outstandingMarketCap"`
		PeLyr                        string `json:"peLyr"`
		PeTtm                        string `json:"peTtm"`
		PbRatio                      string `json:"pbRatio"`
		EpsLyr                       string `json:"epsLyr"`
		EpsTtm                       string `json:"epsTtm"`
		Dividend                     string `json:"dividend"`
		DividendRatio                string `json:"dividendRatio"`
		PriceLimitUp                 string `json:"priceLimitUp"`
		PriceLimitDown               string `json:"priceLimitDown"`
		DividendLfy                  string `json:"dividendLfy"`
		DividendLfyRatio             string `json:"dividendLfyRatio"`
		ExType                       int    `json:"exType"`
		PriceHighestHistory          string `json:"priceHighestHistory"`
		PriceLowestHistory           string `json:"priceLowestHistory"`
		PriceHighest52Week           string `json:"priceHighest_52week"`
		PriceLowest52Week            string `json:"priceLowest_52week"`
		Trust                        struct {
			DividendYield    string `json:"dividendYield"`
			Aum              string `json:"aum"`
			OutstandingUnits string `json:"outstandingUnits"`
			NetAssetValue    string `json:"netAssetValue"`
			AssetClassSc     string `json:"assetClassSc"`
			AssetClassTc     string `json:"assetClassTc"`
			AssetClassEn     string `json:"assetClassEn"`
			AssetClass       int    `json:"assetClass"`
			AssetClassMulti  struct {
				StrContext []interface{} `json:"strContext"`
			} `json:"assetClassMulti"`
		} `json:"trust"`
		PriceBid             string      `json:"priceBid"`
		PriceAsk             string      `json:"priceAsk"`
		VolumeBid            string      `json:"volumeBid"`
		VolumeAsk            string      `json:"volumeAsk"`
		OrderVolumePrecision int         `json:"orderVolumePrecision"`
		HightDirect          string      `json:"hightDirect"`
		OpenDirect           string      `json:"openDirect"`
		LowDirect            string      `json:"lowDirect"`
		Price                string      `json:"price"`
		Change               string      `json:"change"`
		ChangeRatio          string      `json:"changeRatio"`
		PriceDirect          string      `json:"priceDirect"`
		Time                 int64       `json:"time"`
		DelayTime            int         `json:"delayTime"`
		BeforeOpenStockInfo  interface{} `json:"before_open_stock_info"`
		MarketStatus         int         `json:"market_status"`
		MarketStatusText     string      `json:"market_status_text"`
		SparkInfo            struct {
		} `json:"sparkInfo"`
		IsETFTheme bool `json:"isETFTheme"`
	} `json:"stock_info"`
	IndexQuote struct {
		IndexQuotes []struct {
			StockId                      string `json:"stockId"`
			Name                         string `json:"name"`
			MarketType                   int    `json:"marketType"`
			MarketLabel                  string `json:"marketLabel"`
			InstrumentType               int    `json:"instrumentType"`
			StockCode                    string `json:"stockCode"`
			Time                         int64  `json:"time"`
			PriceNominal                 string `json:"priceNominal"`
			PriceLastClose               string `json:"priceLastClose"`
			ServerSendToClientTimeMs     string `json:"serverSendToClientTimeMs"`
			ExchangeDataTimeMs           string `json:"exchangeDataTimeMs"`
			ServerRecvFromExchangeTimeMs string `json:"serverRecvFromExchangeTimeMs"`
			TotalShares                  string `json:"totalShares"`
			TotalMarketCap               string `json:"totalMarketCap"`
			OutstandingShares            string `json:"outstandingShares"`
			OutstandingMarketCap         string `json:"outstandingMarketCap"`
			PriceOpen                    string `json:"priceOpen"`
			PriceHighest                 string `json:"priceHighest"`
			PriceLowest                  string `json:"priceLowest"`
			Volume                       string `json:"volume"`
			VolumeSold                   string `json:"volumeSold"`
			VolumeBought                 string `json:"volumeBought"`
			Turnover                     string `json:"turnover"`
			RatioVolume                  string `json:"ratioVolume"`
			RatioTurnover                string `json:"ratioTurnover"`
			AmplitudePrice               string `json:"amplitudePrice"`
			PriceAverage                 string `json:"priceAverage"`
			ChangeSpeedPrice             string `json:"changeSpeedPrice"`
			RatioBidAsk                  string `json:"ratioBidAsk"`
			VolumePrecision              int    `json:"volumePrecision"`
			OrderVolumePrecision         int    `json:"orderVolumePrecision"`
			Price                        string `json:"price"`
			Change                       string `json:"change"`
			ChangeRatio                  string `json:"changeRatio"`
			PriceDirect                  string `json:"priceDirect"`
			PriceMiddle                  string `json:"priceMiddle"`
			DelayTime                    int    `json:"delayTime"`
			BeforeOpenStockInfo          struct {
				ExchangeTime interface{} `json:"exchange_time"`
				Price        string      `json:"price"`
				Change       string      `json:"change"`
				ChangeRatio  string      `json:"changeRatio"`
				PriceDirect  string      `json:"priceDirect"`
			} `json:"before_open_stock_info"`
		} `json:"indexQuotes"`
	} `json:"index_quote"`
	Financial struct {
		BusinessDataChartList []interface{} `json:"businessDataChartList"`
	} `json:"financial"`
	WidgetSchema struct {
		Schema struct {
			Id            string `json:"id"`
			ComponentName string `json:"componentName"`
			DataConfig    struct {
			} `json:"dataConfig"`
			ComponentConfig struct {
			} `json:"componentConfig"`
			DataSource struct {
			} `json:"dataSource"`
			EnvUserInfo struct {
				IsLogin        bool   `json:"isLogin"`
				UserHomeMarket int    `json:"userHomeMarket"`
				UserStatus     int    `json:"userStatus"`
				PromoteId      int    `json:"promoteId"`
				SubPromoteId   int    `json:"subPromoteId"`
				WidgetId       string `json:"widgetId"`
				BusinessType   string `json:"businessType"`
				BusinessInfo   struct {
					RelatedStockIds []string `json:"relatedStockIds"`
					StyleType       int      `json:"styleType"`
				} `json:"businessInfo"`
				Uid         int    `json:"uid"`
				Lang        int    `json:"lang"`
				UpDownColor int    `json:"upDownColor"`
				Domain      string `json:"domain"`
				Platform    string `json:"platform"`
				IsH5        bool   `json:"isH5"`
				Site        string `json:"site"`
				SchemaId    string `json:"schemaId"`
				ClientIp    string `json:"clientIp"`
				DeviceId    string `json:"deviceId"`
			} `json:"envUserInfo"`
		} `json:"schema"`
	} `json:"widgetSchema"`
	ShowTool    bool `json:"showTool"`
	StockSparks struct {
		List []interface{} `json:"list"`
	} `json:"stock_sparks"`
	PremiumData struct {
		ShowBanner bool   `json:"showBanner"`
		BannerSize string `json:"bannerSize"`
	} `json:"premiumData"`
	StockChartsData struct {
		MinuteChartsData struct {
			StockId string `json:"stockId"`
			List    []struct {
				Time        int     `json:"time"`
				Price       int     `json:"price"`
				CcPrice     float64 `json:"cc_price"`
				Volume      int     `json:"volume"`
				Turnover    int     `json:"turnover"`
				Ratio       string  `json:"ratio"`
				ChangePrice float64 `json:"change_price"`
			} `json:"list"`
			TimeSection []struct {
				Begin int `json:"begin"`
				End   int `json:"end"`
			} `json:"time_section"`
			ServerTime     int `json:"server_time"`
			LastClosePrice int `json:"last_close_price"`
			HighestPrice   int `json:"highest_price"`
			LowestPrice    int `json:"lowest_price"`
		} `json:"minuteChartsData"`
	} `json:"stock_charts_data"`
	StockNews struct {
		ServerTime int64  `json:"server_time"`
		HasMore    bool   `json:"has_more"`
		SeqMark    string `json:"seq_mark"`
		List       []struct {
			Id          int           `json:"id"`
			Title       string        `json:"title"`
			Time        int           `json:"time"`
			Url         string        `json:"url"`
			Source      string        `json:"source"`
			ImptLvl     int           `json:"impt_lvl"`
			ImptTag     string        `json:"impt_tag"`
			ContentTags []interface{} `json:"content_tags"`
			Abstract    string        `json:"abstract"`
			LinkType    int           `json:"link_type"`
		} `json:"list"`
		OverviewList []struct {
			Id          int           `json:"id"`
			Title       string        `json:"title"`
			Time        int           `json:"time"`
			Url         string        `json:"url"`
			Source      string        `json:"source"`
			ImptLvl     int           `json:"impt_lvl"`
			ImptTag     string        `json:"impt_tag"`
			ContentTags []interface{} `json:"content_tags"`
			Abstract    string        `json:"abstract"`
			LinkType    int           `json:"link_type"`
		} `json:"overviewList"`
	} `json:"stock_news"`
	OverView struct {
		FeedData struct {
		} `json:"feedData"`
		SparkList  []interface{} `json:"sparkList"`
		ReqSection int           `json:"reqSection"`
	} `json:"over_view"`
	CommonData struct {
		IsMoomoo bool `json:"isMoomoo"`
		Lang     int  `json:"lang"`
		UserInfo struct {
			Uid            int    `json:"uid"`
			IsLogin        bool   `json:"isLogin"`
			IsMainland     bool   `json:"isMainland"`
			AreaType       string `json:"areaType"`
			CommentNotShow bool   `json:"commentNotShow"`
			IsAsia         bool   `json:"isAsia"`
			MainBroker     int    `json:"mainBroker"`
			UserMarket     string `json:"userMarket"`
		} `json:"userInfo"`
		PreLang     string `json:"preLang"`
		QueryString string `json:"queryString"`
		ShowWatch   bool   `json:"showWatch"`
		NoWatchList bool   `json:"noWatchList"`
		ServerTime  int64  `json:"serverTime"`
		IsMobile    bool   `json:"isMobile"`
		UserData    struct {
		} `json:"userData"`
		Theme    int  `json:"theme"`
		ShowTool bool `json:"showTool"`
		H1Map    struct {
		} `json:"h1Map"`
		Domain        string `json:"domain"`
		Reverse       int    `json:"reverse"`
		AccountStatus bool   `json:"accountStatus"`
	} `json:"common_data"`
	PremiumData1 struct {
		PremiumAccessMap struct {
			Field1 struct {
				HasPermission  bool `json:"hasPermission"`
				ActivationType int  `json:"activationType"`
			} `json:"1"`
			Field2 struct {
				HasPermission  bool `json:"hasPermission"`
				ActivationType int  `json:"activationType"`
			} `json:"19"`
			Field3 struct {
				HasPermission  bool `json:"hasPermission"`
				ActivationType int  `json:"activationType"`
			} `json:"102"`
			Field4 struct {
				HasPermission  bool `json:"hasPermission"`
				ActivationType int  `json:"activationType"`
			} `json:"107"`
			Field5 struct {
				HasPermission  bool `json:"hasPermission"`
				ActivationType int  `json:"activationType"`
			} `json:"108"`
			Field6 struct {
				HasPermission  bool `json:"hasPermission"`
				ActivationType int  `json:"activationType"`
			} `json:"109"`
		} `json:"premiumAccessMap"`
		BannerSize             string `json:"bannerSize"`
		ShowBanner             bool   `json:"showBanner"`
		ShowMiddlePage         bool   `json:"showMiddlePage"`
		MiddlePageAuthorityBit int    `json:"middlePageAuthorityBit"`
		IsPremiumMarket        bool   `json:"isPremiumMarket"`
	} `json:"premium_data"`
	ToolBar struct {
		Label         string `json:"label"`
		IsOpenTool    bool   `json:"isOpenTool"`
		ShallowRefCom struct {
			WATCHLIST string `json:"WATCHLIST"`
			TRADE     string `json:"TRADE"`
		} `json:"shallowRefCom"`
		WatchlistFlash bool `json:"watchlistFlash"`
	} `json:"tool-bar"`
	CountAdd struct {
		Count int `json:"count"`
	} `json:"count_add"`
	WatchList struct {
		ListMap struct {
			Length int `json:"length"`
		} `json:"listMap"`
		IdsMap struct {
		} `json:"idsMap"`
		List     []interface{} `json:"list"`
		Loading  bool          `json:"loading"`
		IsMore   bool          `json:"isMore"`
		Page     int           `json:"page"`
		PageSize int           `json:"page_size"`
	} `json:"watch_list"`
	FixedTop struct {
		StockInfoFixed bool `json:"stockInfoFixed"`
		QuoteFixed     bool `json:"quoteFixed"`
	} `json:"fixed_top"`
	AbTest struct {
		ModalResult struct {
		} `json:"modalResult"`
		BannerShow    bool   `json:"bannerShow"`
		UserMarket    string `json:"userMarket"`
		AccountUrlMap struct {
			Field1 string `json:"0"`
			Field2 string `json:"1"`
			Field3 string `json:"2"`
			Field4 string `json:"3"`
			Field5 string `json:"4"`
			Field6 string `json:"5"`
			Field7 string `json:"6"`
		} `json:"accountUrlMap"`
		AccountStatus bool `json:"accountStatus"`
	} `json:"ab_test"`
	Analysis struct {
		TargetEstimate struct {
			TargetInfo         interface{}   `json:"targetInfo"`
			ClosePriceItemList []interface{} `json:"closePriceItemList"`
		} `json:"targetEstimate"`
		AnalystRating struct {
		} `json:"analystRating"`
	} `json:"analysis"`
	CashFlowIndex struct {
		ListMap struct {
			Field1 []interface{} `json:"0"`
			Field2 []interface{} `json:"1"`
			Field3 []interface{} `json:"2"`
			Field4 []interface{} `json:"3"`
		} `json:"listMap"`
		HasMoreMap struct {
			Field1 bool `json:"0"`
			Field2 bool `json:"1"`
			Field3 bool `json:"2"`
			Field4 bool `json:"3"`
		} `json:"hasMoreMap"`
		FetchId interface{} `json:"fetchId"`
	} `json:"cash_flow_index"`
}
