package types

type EventMean string

const (
	MeanUnkown        EventMean = "MOD_UNKNOWN"
	MeanShotgun       EventMean = "MOD_SHOTGUN"
	MeanGauntlet      EventMean = "MOD_GAUNTLET"
	MeanMachineGun    EventMean = "MOD_MACHINEGUN"
	MeanGrenade       EventMean = "MOD_GRENADE"
	MeanGrenadeSplash EventMean = "MOD_GRENADE_SPLASH"
	MeanRocket        EventMean = "MOD_ROCKET"
	MeanRocketSplash  EventMean = "MOD_ROCKET_SPLASH"
	MeanPlasma        EventMean = "MOD_PLASMA"
	MeanPlasmaSplash  EventMean = "MOD_PLASMA_SPLASH"
	MeanRailgun       EventMean = "MOD_RAILGUN"
	MeanLightning     EventMean = "MOD_LIGHTNING"
	MeanBFG           EventMean = "MOD_BFG"
	MeanBFGSplash     EventMean = "MOD_BFG_SPLASH"
	MeanWater         EventMean = "MOD_WATER"
	MeanSlime         EventMean = "MOD_SLIME"
	MeanLava          EventMean = "MOD_LAVA"
	MeanCrush         EventMean = "MOD_CRUSH"
	MeanTelefrag      EventMean = "MOD_TELEFRAG"
	MeanFalling       EventMean = "MOD_FALLING"
	MeanSuicide       EventMean = "MOD_SUICIDE"
	MeanTargetLaser   EventMean = "MOD_TARGET_LASER"
	MeanTriggerHurt   EventMean = "MOD_TRIGGER_HURT"
	MeanNail          EventMean = "MOD_NAIL"
	MeanChaingun      EventMean = "MOD_CHAINGUN"
	MeanProximityMine EventMean = "MOD_PROXIMITY_MINE"
	MeanKamikaze      EventMean = "MOD_KAMIKAZE"
	MeanJuiced        EventMean = "MOD_JUICED"
	MeanGrapple       EventMean = "MOD_GRAPPLE"
)

func EventMeans() []EventMean {
	return []EventMean{
		MeanUnkown,
		MeanShotgun,
		MeanGauntlet,
		MeanMachineGun,
		MeanGrenade,
		MeanGrenadeSplash,
		MeanRocket,
		MeanRocketSplash,
		MeanPlasma,
		MeanPlasmaSplash,
		MeanRailgun,
		MeanLightning,
		MeanBFG,
		MeanBFGSplash,
		MeanWater,
		MeanSlime,
		MeanLava,
		MeanCrush,
		MeanTelefrag,
		MeanFalling,
		MeanSuicide,
		MeanTargetLaser,
		MeanTriggerHurt,
		MeanNail,
		MeanChaingun,
		MeanProximityMine,
		MeanKamikaze,
		MeanJuiced,
		MeanGrapple,
	}
}
