#ifndef __BASIC_HLSL__
#define __BASIC_HLSL__



// note: here is the hlsl version, only included in hlsl

// MyTrickyFun this is the GPU version of the tricky function
float MyTrickyFun(float x) {
	return 16; // ok actually not tricky here, but whatever
}


// FastExp is a quartic spline approximation to the Exp function, by N.N. Schraudolph
// It does not have any of the sanity checking of a standard method -- returns
// nonsense when arg is out of range.  Runs in 2.23ns vs. 6.3ns for 64bit which is faster
// than exp actually.
float FastExp(float x) {
	if (x <= -88.76731) { // this doesn't add anything and -exp is main use-case anyway
		return 0;
	}
	int i = int(12102203*x) + 127*(1<<23);
	int m = i >> 7 & 0xFFFF; // copy mantissa
	i += (((((((((((3537 * m) >> 16) + 13668) * m) >> 18) + 15817) * m) >> 14) - 80470) * m) >> 11);
	return asfloat(uint(i));
}

// NeuronFlags are bit-flags encoding relevant binary state for neurons
typedef int NeuronFlags;

// The neuron flags

// NeuronOff flag indicates that this neuron has been turned off (i.e., lesioned)
static const NeuronFlags NeuronOff = 1;

// NeuronHasExt means the neuron has external input in its Ext field
static const NeuronFlags NeuronHasExt = 1 << 2;

// NeuronHasTarg means the neuron has external target input in its Target field
static const NeuronFlags NeuronHasTarg = 1 << 3;

// NeuronHasCmpr means the neuron has external comparison input in its Target field -- used for computing
// comparison statistics but does not drive neural activity ever
static const NeuronFlags NeuronHasCmpr = 1 << 4;

// Modes are evaluation modes (Training, Testing, etc)
typedef int Modes;

// The evaluation modes

static const Modes NoEvalMode = 0;

// AllModes indicates that the log should occur over all modes present in other items.
static const Modes AllModes = 1;

// Train is this a training mode for the env
static const Modes Train = 2;

// Test is this a test mode for the env
static const Modes Test = 3;

// DataStruct has the test data
struct DataStruct {

	// raw value
	float Raw;

	// integrated value
	float Integ;

	// exp of integ
	float Exp;

	// must pad to multiple of 4 floats for arrays
	float Pad2;
};

// ParamStruct has the test params
struct ParamStruct {

	// rate constant in msec
	float Tau;

	// 1/Tau
	float     Dt;
	int Option; // note: standard bool doesn't work

	float pad; // comment this out to trigger alignment warning
	void IntegFromRaw(inout DataStruct ds, inout float modArg) {
		// note: the following are just to test basic control structures
		float newVal = this.Dt*(ds.Raw-ds.Integ) + modArg;
		if (newVal < -10 || this.Option==1) {
			newVal = -10;
		}
		ds.Integ += newVal;
		ds.Exp = exp(-ds.Integ);
	}

	void AnotherMeth(inout DataStruct ds) {
		for (int i = 0; i < 10; i++) {
			ds.Integ *= 0.99;
		}
		NeuronFlags flag;
		flag &=~NeuronHasExt; // clear flag -- op doesn't exist in C

		Modes mode = Test;
		switch (mode) {
		case 3:
		// fallthrough

		case 2:{
			float ab = float(.5);
			ds.Exp *= ab;
			break; }
		default:{
			float ab = float(1);
			ds.Exp *= ab;
			break; }
		}
	}

};


[[vk::binding(0, 0)]] StructuredBuffer<ParamStruct> Params;
[[vk::binding(0, 1)]] RWStructuredBuffer<DataStruct> Data;
[numthreads(1, 1, 1)]
void main(uint3 idx : SV_DispatchThreadID) {
    Params[0].IntegFromRaw(Data[idx.x], Data[idx.x].Pad2);
}
#endif // __BASIC_HLSL__
