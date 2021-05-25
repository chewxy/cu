package cudnn

/* Generated by gencudnn. DO NOT EDIT */

// #include <cudnn.h>
import "C"

type ActivationMode int

const (
	Sigmoid     ActivationMode = C.CUDNN_ACTIVATION_SIGMOID
	ReLU        ActivationMode = C.CUDNN_ACTIVATION_RELU
	Tanh        ActivationMode = C.CUDNN_ACTIVATION_TANH
	ClippedReLU ActivationMode = C.CUDNN_ACTIVATION_CLIPPED_RELU
	Elu         ActivationMode = C.CUDNN_ACTIVATION_ELU
	Identity    ActivationMode = C.CUDNN_ACTIVATION_IDENTITY
)

// C returns the C representation of ActivationMode
func (e ActivationMode) C() C.cudnnActivationMode_t { return C.cudnnActivationMode_t(e) }

type BackendAttributeName int

const (
	PointwiseMode                         BackendAttributeName = C.CUDNN_ATTR_POINTWISE_MODE
	PointwiseMathPrec                     BackendAttributeName = C.CUDNN_ATTR_POINTWISE_MATH_PREC
	PointwiseNanPropagation               BackendAttributeName = C.CUDNN_ATTR_POINTWISE_NAN_PROPAGATION
	PointwiseReluLowerClip                BackendAttributeName = C.CUDNN_ATTR_POINTWISE_RELU_LOWER_CLIP
	PointwiseReluUpperClip                BackendAttributeName = C.CUDNN_ATTR_POINTWISE_RELU_UPPER_CLIP
	ConvolutionCompType                   BackendAttributeName = C.CUDNN_ATTR_CONVOLUTION_COMP_TYPE
	ConvolutionConvMode                   BackendAttributeName = C.CUDNN_ATTR_CONVOLUTION_CONV_MODE
	ConvolutionDilations                  BackendAttributeName = C.CUDNN_ATTR_CONVOLUTION_DILATIONS
	ConvolutionFilterStrides              BackendAttributeName = C.CUDNN_ATTR_CONVOLUTION_FILTER_STRIDES
	ConvolutionPostPaddings               BackendAttributeName = C.CUDNN_ATTR_CONVOLUTION_POST_PADDINGS
	ConvolutionPrePaddings                BackendAttributeName = C.CUDNN_ATTR_CONVOLUTION_PRE_PADDINGS
	ConvolutionSpatialDims                BackendAttributeName = C.CUDNN_ATTR_CONVOLUTION_SPATIAL_DIMS
	EngineheurMode                        BackendAttributeName = C.CUDNN_ATTR_ENGINEHEUR_MODE
	EngineheurOperationGraph              BackendAttributeName = C.CUDNN_ATTR_ENGINEHEUR_OPERATION_GRAPH
	EngineheurResults                     BackendAttributeName = C.CUDNN_ATTR_ENGINEHEUR_RESULTS
	EnginecfgEngine                       BackendAttributeName = C.CUDNN_ATTR_ENGINECFG_ENGINE
	EnginecfgIntermediateInfo             BackendAttributeName = C.CUDNN_ATTR_ENGINECFG_INTERMEDIATE_INFO
	EnginecfgKnobChoices                  BackendAttributeName = C.CUDNN_ATTR_ENGINECFG_KNOB_CHOICES
	ExecutionPlanHandle                   BackendAttributeName = C.CUDNN_ATTR_EXECUTION_PLAN_HANDLE
	ExecutionPlanEngineConfig             BackendAttributeName = C.CUDNN_ATTR_EXECUTION_PLAN_ENGINE_CONFIG
	ExecutionPlanWorkspaceSize            BackendAttributeName = C.CUDNN_ATTR_EXECUTION_PLAN_WORKSPACE_SIZE
	ExecutionPlanComputedIntermediateUids BackendAttributeName = C.CUDNN_ATTR_EXECUTION_PLAN_COMPUTED_INTERMEDIATE_UIDS
	ExecutionPlanRunOnlyIntermediateUids  BackendAttributeName = C.CUDNN_ATTR_EXECUTION_PLAN_RUN_ONLY_INTERMEDIATE_UIDS
	IntermediateInfoUniqueId              BackendAttributeName = C.CUDNN_ATTR_INTERMEDIATE_INFO_UNIQUE_ID
	IntermediateInfoSize                  BackendAttributeName = C.CUDNN_ATTR_INTERMEDIATE_INFO_SIZE
	IntermediateInfoDependentDataUids     BackendAttributeName = C.CUDNN_ATTR_INTERMEDIATE_INFO_DEPENDENT_DATA_UIDS
	IntermediateInfoDependentAttributes   BackendAttributeName = C.CUDNN_ATTR_INTERMEDIATE_INFO_DEPENDENT_ATTRIBUTES
	KnobChoiceKnobType                    BackendAttributeName = C.CUDNN_ATTR_KNOB_CHOICE_KNOB_TYPE
	KnobChoiceKnobValue                   BackendAttributeName = C.CUDNN_ATTR_KNOB_CHOICE_KNOB_VALUE
	OperationConvolutionForwardAlpha      BackendAttributeName = C.CUDNN_ATTR_OPERATION_CONVOLUTION_FORWARD_ALPHA
	OperationConvolutionForwardBeta       BackendAttributeName = C.CUDNN_ATTR_OPERATION_CONVOLUTION_FORWARD_BETA
	OperationConvolutionForwardConvDesc   BackendAttributeName = C.CUDNN_ATTR_OPERATION_CONVOLUTION_FORWARD_CONV_DESC
	OperationConvolutionForwardW          BackendAttributeName = C.CUDNN_ATTR_OPERATION_CONVOLUTION_FORWARD_W
	OperationConvolutionForwardX          BackendAttributeName = C.CUDNN_ATTR_OPERATION_CONVOLUTION_FORWARD_X
	OperationConvolutionForwardY          BackendAttributeName = C.CUDNN_ATTR_OPERATION_CONVOLUTION_FORWARD_Y
	OperationConvolutionBwdDataAlpha      BackendAttributeName = C.CUDNN_ATTR_OPERATION_CONVOLUTION_BWD_DATA_ALPHA
	OperationConvolutionBwdDataBeta       BackendAttributeName = C.CUDNN_ATTR_OPERATION_CONVOLUTION_BWD_DATA_BETA
	OperationConvolutionBwdDataConvDesc   BackendAttributeName = C.CUDNN_ATTR_OPERATION_CONVOLUTION_BWD_DATA_CONV_DESC
	OperationConvolutionBwdDataW          BackendAttributeName = C.CUDNN_ATTR_OPERATION_CONVOLUTION_BWD_DATA_W
	OperationConvolutionBwdDataDx         BackendAttributeName = C.CUDNN_ATTR_OPERATION_CONVOLUTION_BWD_DATA_DX
	OperationConvolutionBwdDataDy         BackendAttributeName = C.CUDNN_ATTR_OPERATION_CONVOLUTION_BWD_DATA_DY
	OperationConvolutionBwdFilterAlpha    BackendAttributeName = C.CUDNN_ATTR_OPERATION_CONVOLUTION_BWD_FILTER_ALPHA
	OperationConvolutionBwdFilterBeta     BackendAttributeName = C.CUDNN_ATTR_OPERATION_CONVOLUTION_BWD_FILTER_BETA
	OperationConvolutionBwdFilterConvDesc BackendAttributeName = C.CUDNN_ATTR_OPERATION_CONVOLUTION_BWD_FILTER_CONV_DESC
	OperationConvolutionBwdFilterDw       BackendAttributeName = C.CUDNN_ATTR_OPERATION_CONVOLUTION_BWD_FILTER_DW
	OperationConvolutionBwdFilterX        BackendAttributeName = C.CUDNN_ATTR_OPERATION_CONVOLUTION_BWD_FILTER_X
	OperationConvolutionBwdFilterDy       BackendAttributeName = C.CUDNN_ATTR_OPERATION_CONVOLUTION_BWD_FILTER_DY
	OperationPointwisePwDescriptor        BackendAttributeName = C.CUDNN_ATTR_OPERATION_POINTWISE_PW_DESCRIPTOR
	OperationPointwiseXdesc               BackendAttributeName = C.CUDNN_ATTR_OPERATION_POINTWISE_XDESC
	OperationPointwiseBdesc               BackendAttributeName = C.CUDNN_ATTR_OPERATION_POINTWISE_BDESC
	OperationPointwiseYdesc               BackendAttributeName = C.CUDNN_ATTR_OPERATION_POINTWISE_YDESC
	OperationPointwiseAlpha1              BackendAttributeName = C.CUDNN_ATTR_OPERATION_POINTWISE_ALPHA1
	OperationPointwiseAlpha2              BackendAttributeName = C.CUDNN_ATTR_OPERATION_POINTWISE_ALPHA2
	OperationGenstatsMode                 BackendAttributeName = C.CUDNN_ATTR_OPERATION_GENSTATS_MODE
	OperationGenstatsMathPrec             BackendAttributeName = C.CUDNN_ATTR_OPERATION_GENSTATS_MATH_PREC
	OperationGenstatsXdesc                BackendAttributeName = C.CUDNN_ATTR_OPERATION_GENSTATS_XDESC
	OperationGenstatsSumdesc              BackendAttributeName = C.CUDNN_ATTR_OPERATION_GENSTATS_SUMDESC
	OperationGenstatsSqsumdesc            BackendAttributeName = C.CUDNN_ATTR_OPERATION_GENSTATS_SQSUMDESC
	OperationgraphHandle                  BackendAttributeName = C.CUDNN_ATTR_OPERATIONGRAPH_HANDLE
	OperationgraphOps                     BackendAttributeName = C.CUDNN_ATTR_OPERATIONGRAPH_OPS
	OperationgraphEngineGlobalCount       BackendAttributeName = C.CUDNN_ATTR_OPERATIONGRAPH_ENGINE_GLOBAL_COUNT
	TensorByteAlignment                   BackendAttributeName = C.CUDNN_ATTR_TENSOR_BYTE_ALIGNMENT
	TensorDataType                        BackendAttributeName = C.CUDNN_ATTR_TENSOR_DATA_TYPE
	TensorDimensions                      BackendAttributeName = C.CUDNN_ATTR_TENSOR_DIMENSIONS
	TensorStrides                         BackendAttributeName = C.CUDNN_ATTR_TENSOR_STRIDES
	TensorVectorCount                     BackendAttributeName = C.CUDNN_ATTR_TENSOR_VECTOR_COUNT
	TensorVectorizedDimension             BackendAttributeName = C.CUDNN_ATTR_TENSOR_VECTORIZED_DIMENSION
	TensorUniqueId                        BackendAttributeName = C.CUDNN_ATTR_TENSOR_UNIQUE_ID
	TensorIsVirtual                       BackendAttributeName = C.CUDNN_ATTR_TENSOR_IS_VIRTUAL
	VariantPackUniqueIds                  BackendAttributeName = C.CUDNN_ATTR_VARIANT_PACK_UNIQUE_IDS
	VariantPackDataPointers               BackendAttributeName = C.CUDNN_ATTR_VARIANT_PACK_DATA_POINTERS
	VariantPackIntermediates              BackendAttributeName = C.CUDNN_ATTR_VARIANT_PACK_INTERMEDIATES
	VariantPackWorkspace                  BackendAttributeName = C.CUDNN_ATTR_VARIANT_PACK_WORKSPACE
	LayoutInfoTensorUid                   BackendAttributeName = C.CUDNN_ATTR_LAYOUT_INFO_TENSOR_UID
	LayoutInfoTypes                       BackendAttributeName = C.CUDNN_ATTR_LAYOUT_INFO_TYPES
	KnobInfoType                          BackendAttributeName = C.CUDNN_ATTR_KNOB_INFO_TYPE
	KnobInfoMaximumValue                  BackendAttributeName = C.CUDNN_ATTR_KNOB_INFO_MAXIMUM_VALUE
	KnobInfoMinimumValue                  BackendAttributeName = C.CUDNN_ATTR_KNOB_INFO_MINIMUM_VALUE
	KnobInfoStride                        BackendAttributeName = C.CUDNN_ATTR_KNOB_INFO_STRIDE
	EngineOperationGraph                  BackendAttributeName = C.CUDNN_ATTR_ENGINE_OPERATION_GRAPH
	EngineGlobalIndex                     BackendAttributeName = C.CUDNN_ATTR_ENGINE_GLOBAL_INDEX
	EngineKnobInfo                        BackendAttributeName = C.CUDNN_ATTR_ENGINE_KNOB_INFO
	EngineNumericalNote                   BackendAttributeName = C.CUDNN_ATTR_ENGINE_NUMERICAL_NOTE
	EngineLayoutInfo                      BackendAttributeName = C.CUDNN_ATTR_ENGINE_LAYOUT_INFO
)

// C returns the C representation of BackendAttributeName
func (e BackendAttributeName) C() C.cudnnBackendAttributeName_t {
	return C.cudnnBackendAttributeName_t(e)
}

type BackendAttributeType int

const (
	BackendAttrHandle            BackendAttributeType = C.CUDNN_TYPE_HANDLE
	BackendAttrDataType          BackendAttributeType = C.CUDNN_TYPE_DATA_TYPE
	BackendAttrBoolean           BackendAttributeType = C.CUDNN_TYPE_BOOLEAN
	BackendAttrInt64             BackendAttributeType = C.CUDNN_TYPE_INT64
	BackendAttrFloat             BackendAttributeType = C.CUDNN_TYPE_FLOAT
	BackendAttrDouble            BackendAttributeType = C.CUDNN_TYPE_DOUBLE
	BackendAttrVoidPtr           BackendAttributeType = C.CUDNN_TYPE_VOID_PTR
	BackendAttrConvolutionMode   BackendAttributeType = C.CUDNN_TYPE_CONVOLUTION_MODE
	BackendAttrHeurMode          BackendAttributeType = C.CUDNN_TYPE_HEUR_MODE
	BackendAttrKnobType          BackendAttributeType = C.CUDNN_TYPE_KNOB_TYPE
	BackendAttrNanPropogation    BackendAttributeType = C.CUDNN_TYPE_NAN_PROPOGATION
	BackendAttrNumericalNote     BackendAttributeType = C.CUDNN_TYPE_NUMERICAL_NOTE
	BackendAttrLayoutType        BackendAttributeType = C.CUDNN_TYPE_LAYOUT_TYPE
	BackendAttrAttribName        BackendAttributeType = C.CUDNN_TYPE_ATTRIB_NAME
	BackendAttrPointwiseMode     BackendAttributeType = C.CUDNN_TYPE_POINTWISE_MODE
	BackendAttrBackendDescriptor BackendAttributeType = C.CUDNN_TYPE_BACKEND_DESCRIPTOR
	BackendAttrGenstatsMode      BackendAttributeType = C.CUDNN_TYPE_GENSTATS_MODE
)

// C returns the C representation of BackendAttributeType
func (e BackendAttributeType) C() C.cudnnBackendAttributeType_t {
	return C.cudnnBackendAttributeType_t(e)
}

type BackendDescriptorType int

const (
	PointwiseDescriptor                          BackendDescriptorType = C.CUDNN_BACKEND_POINTWISE_DESCRIPTOR
	ConvolutionDescriptor                        BackendDescriptorType = C.CUDNN_BACKEND_CONVOLUTION_DESCRIPTOR
	EngineDescriptor                             BackendDescriptorType = C.CUDNN_BACKEND_ENGINE_DESCRIPTOR
	EnginecfgDescriptor                          BackendDescriptorType = C.CUDNN_BACKEND_ENGINECFG_DESCRIPTOR
	EngineheurDescriptor                         BackendDescriptorType = C.CUDNN_BACKEND_ENGINEHEUR_DESCRIPTOR
	ExecutionPlanDescriptor                      BackendDescriptorType = C.CUDNN_BACKEND_EXECUTION_PLAN_DESCRIPTOR
	IntermediateInfoDescriptor                   BackendDescriptorType = C.CUDNN_BACKEND_INTERMEDIATE_INFO_DESCRIPTOR
	KnobChoiceDescriptor                         BackendDescriptorType = C.CUDNN_BACKEND_KNOB_CHOICE_DESCRIPTOR
	KnobInfoDescriptor                           BackendDescriptorType = C.CUDNN_BACKEND_KNOB_INFO_DESCRIPTOR
	LayoutInfoDescriptor                         BackendDescriptorType = C.CUDNN_BACKEND_LAYOUT_INFO_DESCRIPTOR
	OperationConvolutionForwardDescriptor        BackendDescriptorType = C.CUDNN_BACKEND_OPERATION_CONVOLUTION_FORWARD_DESCRIPTOR
	OperationConvolutionBackwardFilterDescriptor BackendDescriptorType = C.CUDNN_BACKEND_OPERATION_CONVOLUTION_BACKWARD_FILTER_DESCRIPTOR
	OperationConvolutionBackwardDataDescriptor   BackendDescriptorType = C.CUDNN_BACKEND_OPERATION_CONVOLUTION_BACKWARD_DATA_DESCRIPTOR
	OperationPointwiseDescriptor                 BackendDescriptorType = C.CUDNN_BACKEND_OPERATION_POINTWISE_DESCRIPTOR
	OperationGenStatsDescriptor                  BackendDescriptorType = C.CUDNN_BACKEND_OPERATION_GEN_STATS_DESCRIPTOR
	OperationgraphDescriptor                     BackendDescriptorType = C.CUDNN_BACKEND_OPERATIONGRAPH_DESCRIPTOR
	VariantPackDescriptor                        BackendDescriptorType = C.CUDNN_BACKEND_VARIANT_PACK_DESCRIPTOR
	TensorDescriptor                             BackendDescriptorType = C.CUDNN_BACKEND_TENSOR_DESCRIPTOR
)

// C returns the C representation of BackendDescriptorType
func (e BackendDescriptorType) C() C.cudnnBackendDescriptorType_t {
	return C.cudnnBackendDescriptorType_t(e)
}

type BackendHeurMode int

const (
	Instant BackendHeurMode = C.CUDNN_HEUR_MODE_INSTANT
	SCount  BackendHeurMode = C.CUDNN_HEUR_MODES_COUNT
)

// C returns the C representation of BackendHeurMode
func (e BackendHeurMode) C() C.cudnnBackendHeurMode_t { return C.cudnnBackendHeurMode_t(e) }

type BackendKnobType int

const (
	SplitK        BackendKnobType = C.CUDNN_KNOB_TYPE_SPLIT_K
	Swizzle       BackendKnobType = C.CUDNN_KNOB_TYPE_SWIZZLE
	TileSize      BackendKnobType = C.CUDNN_KNOB_TYPE_TILE_SIZE
	UseTex        BackendKnobType = C.CUDNN_KNOB_TYPE_USE_TEX
	Edge          BackendKnobType = C.CUDNN_KNOB_TYPE_EDGE
	Kblock        BackendKnobType = C.CUDNN_KNOB_TYPE_KBLOCK
	Ldga          BackendKnobType = C.CUDNN_KNOB_TYPE_LDGA
	Ldgb          BackendKnobType = C.CUDNN_KNOB_TYPE_LDGB
	ChunkK        BackendKnobType = C.CUDNN_KNOB_TYPE_CHUNK_K
	SplitH        BackendKnobType = C.CUDNN_KNOB_TYPE_SPLIT_H
	WinoTile      BackendKnobType = C.CUDNN_KNOB_TYPE_WINO_TILE
	Multiply      BackendKnobType = C.CUDNN_KNOB_TYPE_MULTIPLY
	SplitKBuf     BackendKnobType = C.CUDNN_KNOB_TYPE_SPLIT_K_BUF
	Tilek         BackendKnobType = C.CUDNN_KNOB_TYPE_TILEK
	Stages        BackendKnobType = C.CUDNN_KNOB_TYPE_STAGES
	ReductionMode BackendKnobType = C.CUDNN_KNOB_TYPE_REDUCTION_MODE
	CtaSplitKMode BackendKnobType = C.CUDNN_KNOB_TYPE_CTA_SPLIT_K_MODE
	SplitKSlc     BackendKnobType = C.CUDNN_KNOB_TYPE_SPLIT_K_SLC
	IdxMode       BackendKnobType = C.CUDNN_KNOB_TYPE_IDX_MODE
	Sliced        BackendKnobType = C.CUDNN_KNOB_TYPE_SLICED
	SplitRs       BackendKnobType = C.CUDNN_KNOB_TYPE_SPLIT_RS
	Singlebuffer  BackendKnobType = C.CUDNN_KNOB_TYPE_SINGLEBUFFER
	Ldgc          BackendKnobType = C.CUDNN_KNOB_TYPE_LDGC
	Specfilt      BackendKnobType = C.CUDNN_KNOB_TYPE_SPECFILT
	Counts        BackendKnobType = C.CUDNN_KNOB_TYPE_COUNTS
)

// C returns the C representation of BackendKnobType
func (e BackendKnobType) C() C.cudnnBackendKnobType_t { return C.cudnnBackendKnobType_t(e) }

type BackendLayoutType int

const (
	PreferredNchw   BackendLayoutType = C.CUDNN_LAYOUT_TYPE_PREFERRED_NCHW
	PreferredNhwc   BackendLayoutType = C.CUDNN_LAYOUT_TYPE_PREFERRED_NHWC
	PreferredPad4ck BackendLayoutType = C.CUDNN_LAYOUT_TYPE_PREFERRED_PAD4CK
	PreferredPad8ck BackendLayoutType = C.CUDNN_LAYOUT_TYPE_PREFERRED_PAD8CK
	Count           BackendLayoutType = C.CUDNN_LAYOUT_TYPE_COUNT
)

// C returns the C representation of BackendLayoutType
func (e BackendLayoutType) C() C.cudnnBackendLayoutType_t { return C.cudnnBackendLayoutType_t(e) }

type BackendNumericalNote int

const (
	TensorCore                BackendNumericalNote = C.CUDNN_NUMERICAL_NOTE_TENSOR_CORE
	DownConvertInputs         BackendNumericalNote = C.CUDNN_NUMERICAL_NOTE_DOWN_CONVERT_INPUTS
	ReducedPrecisionReduction BackendNumericalNote = C.CUDNN_NUMERICAL_NOTE_REDUCED_PRECISION_REDUCTION
	Fft                       BackendNumericalNote = C.CUDNN_NUMERICAL_NOTE_FFT
	Nondeterministic          BackendNumericalNote = C.CUDNN_NUMERICAL_NOTE_NONDETERMINISTIC
	Winograd                  BackendNumericalNote = C.CUDNN_NUMERICAL_NOTE_WINOGRAD
	TypeCount                 BackendNumericalNote = C.CUDNN_NUMERICAL_NOTE_TYPE_COUNT
)

// C returns the C representation of BackendNumericalNote
func (e BackendNumericalNote) C() C.cudnnBackendNumericalNote_t {
	return C.cudnnBackendNumericalNote_t(e)
}

type BatchNormMode int

const (
	PerActivation     BatchNormMode = C.CUDNN_BATCHNORM_PER_ACTIVATION
	Spatial           BatchNormMode = C.CUDNN_BATCHNORM_SPATIAL
	SpatialPersistent BatchNormMode = C.CUDNN_BATCHNORM_SPATIAL_PERSISTENT
)

// C returns the C representation of BatchNormMode
func (e BatchNormMode) C() C.cudnnBatchNormMode_t { return C.cudnnBatchNormMode_t(e) }

type BatchNormOps int

const (
	BatchNorm     BatchNormOps = C.CUDNN_BATCHNORM_OPS_BN
	Activation    BatchNormOps = C.CUDNN_BATCHNORM_OPS_BN_ACTIVATION
	AddActivation BatchNormOps = C.CUDNN_BATCHNORM_OPS_BN_ADD_ACTIVATION
)

// C returns the C representation of BatchNormOps
func (e BatchNormOps) C() C.cudnnBatchNormOps_t { return C.cudnnBatchNormOps_t(e) }

type CTCLossAlgo int

const (
	DeterministicCTCLoss    CTCLossAlgo = C.CUDNN_CTC_LOSS_ALGO_DETERMINISTIC
	NonDeterministicCTCLoss CTCLossAlgo = C.CUDNN_CTC_LOSS_ALGO_NON_DETERMINISTIC
)

// C returns the C representation of CTCLossAlgo
func (e CTCLossAlgo) C() C.cudnnCTCLossAlgo_t { return C.cudnnCTCLossAlgo_t(e) }

type DataType int

const (
	Float   DataType = C.CUDNN_DATA_FLOAT
	Double  DataType = C.CUDNN_DATA_DOUBLE
	Half    DataType = C.CUDNN_DATA_HALF
	Int8    DataType = C.CUDNN_DATA_INT8
	Int32   DataType = C.CUDNN_DATA_INT32
	Int8x4  DataType = C.CUDNN_DATA_INT8x4
	Uint8   DataType = C.CUDNN_DATA_UINT8
	Uint8x4 DataType = C.CUDNN_DATA_UINT8x4
	Int8x32 DataType = C.CUDNN_DATA_INT8x32
)

// C returns the C representation of DataType
func (e DataType) C() C.cudnnDataType_t { return C.cudnnDataType_t(e) }

type Determinism int

const (
	NonDeterministic Determinism = C.CUDNN_NON_DETERMINISTIC
	Deterministic    Determinism = C.CUDNN_DETERMINISTIC
)

// C returns the C representation of Determinism
func (e Determinism) C() C.cudnnDeterminism_t { return C.cudnnDeterminism_t(e) }

type DirectionMode int

const (
	Unidirectional DirectionMode = C.CUDNN_UNIDIRECTIONAL
	Bidirectional  DirectionMode = C.CUDNN_BIDIRECTIONAL
)

// C returns the C representation of DirectionMode
func (e DirectionMode) C() C.cudnnDirectionMode_t { return C.cudnnDirectionMode_t(e) }

type DivNormMode int

const (
	PrecomputedMeans DivNormMode = C.CUDNN_DIVNORM_PRECOMPUTED_MEANS
)

// C returns the C representation of DivNormMode
func (e DivNormMode) C() C.cudnnDivNormMode_t { return C.cudnnDivNormMode_t(e) }

type ErrQueryMode int

const (
	Rawcode     ErrQueryMode = C.CUDNN_ERRQUERY_RAWCODE
	Nonblocking ErrQueryMode = C.CUDNN_ERRQUERY_NONBLOCKING
	Blocking    ErrQueryMode = C.CUDNN_ERRQUERY_BLOCKING
)

// C returns the C representation of ErrQueryMode
func (e ErrQueryMode) C() C.cudnnErrQueryMode_t { return C.cudnnErrQueryMode_t(e) }

type FoldingDirection int

const (
	Fold   FoldingDirection = C.CUDNN_TRANSFORM_FOLD
	Unfold FoldingDirection = C.CUDNN_TRANSFORM_UNFOLD
)

// C returns the C representation of FoldingDirection
func (e FoldingDirection) C() C.cudnnFoldingDirection_t { return C.cudnnFoldingDirection_t(e) }

type ForwardMode int

const (
	Inference ForwardMode = C.CUDNN_FWD_MODE_INFERENCE
	Training  ForwardMode = C.CUDNN_FWD_MODE_TRAINING
)

// C returns the C representation of ForwardMode
func (e ForwardMode) C() C.cudnnForwardMode_t { return C.cudnnForwardMode_t(e) }

type FusedOpsConstParamLabel int

const (
	Xdesc                        FusedOpsConstParamLabel = C.CUDNN_PARAM_XDESC
	XdataPlaceholder             FusedOpsConstParamLabel = C.CUDNN_PARAM_XDATA_PLACEHOLDER
	BnMode                       FusedOpsConstParamLabel = C.CUDNN_PARAM_BN_MODE
	BnEqscalebiasDesc            FusedOpsConstParamLabel = C.CUDNN_PARAM_BN_EQSCALEBIAS_DESC
	BnEqscalePlaceholder         FusedOpsConstParamLabel = C.CUDNN_PARAM_BN_EQSCALE_PLACEHOLDER
	BnEqbiasPlaceholder          FusedOpsConstParamLabel = C.CUDNN_PARAM_BN_EQBIAS_PLACEHOLDER
	ActivationDesc               FusedOpsConstParamLabel = C.CUDNN_PARAM_ACTIVATION_DESC
	ConvDesc                     FusedOpsConstParamLabel = C.CUDNN_PARAM_CONV_DESC
	Wdesc                        FusedOpsConstParamLabel = C.CUDNN_PARAM_WDESC
	WdataPlaceholder             FusedOpsConstParamLabel = C.CUDNN_PARAM_WDATA_PLACEHOLDER
	Dwdesc                       FusedOpsConstParamLabel = C.CUDNN_PARAM_DWDESC
	DwdataPlaceholder            FusedOpsConstParamLabel = C.CUDNN_PARAM_DWDATA_PLACEHOLDER
	Ydesc                        FusedOpsConstParamLabel = C.CUDNN_PARAM_YDESC
	YdataPlaceholder             FusedOpsConstParamLabel = C.CUDNN_PARAM_YDATA_PLACEHOLDER
	Dydesc                       FusedOpsConstParamLabel = C.CUDNN_PARAM_DYDESC
	DydataPlaceholder            FusedOpsConstParamLabel = C.CUDNN_PARAM_DYDATA_PLACEHOLDER
	YstatsDesc                   FusedOpsConstParamLabel = C.CUDNN_PARAM_YSTATS_DESC
	YsumPlaceholder              FusedOpsConstParamLabel = C.CUDNN_PARAM_YSUM_PLACEHOLDER
	YsqsumPlaceholder            FusedOpsConstParamLabel = C.CUDNN_PARAM_YSQSUM_PLACEHOLDER
	BnScalebiasMeanvarDesc       FusedOpsConstParamLabel = C.CUDNN_PARAM_BN_SCALEBIAS_MEANVAR_DESC
	BnScalePlaceholder           FusedOpsConstParamLabel = C.CUDNN_PARAM_BN_SCALE_PLACEHOLDER
	BnBiasPlaceholder            FusedOpsConstParamLabel = C.CUDNN_PARAM_BN_BIAS_PLACEHOLDER
	BnSavedMeanPlaceholder       FusedOpsConstParamLabel = C.CUDNN_PARAM_BN_SAVED_MEAN_PLACEHOLDER
	BnSavedInvstdPlaceholder     FusedOpsConstParamLabel = C.CUDNN_PARAM_BN_SAVED_INVSTD_PLACEHOLDER
	BnRunningMeanPlaceholder     FusedOpsConstParamLabel = C.CUDNN_PARAM_BN_RUNNING_MEAN_PLACEHOLDER
	BnRunningVarPlaceholder      FusedOpsConstParamLabel = C.CUDNN_PARAM_BN_RUNNING_VAR_PLACEHOLDER
	Zdesc                        FusedOpsConstParamLabel = C.CUDNN_PARAM_ZDESC
	ZdataPlaceholder             FusedOpsConstParamLabel = C.CUDNN_PARAM_ZDATA_PLACEHOLDER
	BnZEqscalebiasDesc           FusedOpsConstParamLabel = C.CUDNN_PARAM_BN_Z_EQSCALEBIAS_DESC
	BnZEqscalePlaceholder        FusedOpsConstParamLabel = C.CUDNN_PARAM_BN_Z_EQSCALE_PLACEHOLDER
	BnZEqbiasPlaceholder         FusedOpsConstParamLabel = C.CUDNN_PARAM_BN_Z_EQBIAS_PLACEHOLDER
	ActivationBitmaskDesc        FusedOpsConstParamLabel = C.CUDNN_PARAM_ACTIVATION_BITMASK_DESC
	ActivationBitmaskPlaceholder FusedOpsConstParamLabel = C.CUDNN_PARAM_ACTIVATION_BITMASK_PLACEHOLDER
	Dxdesc                       FusedOpsConstParamLabel = C.CUDNN_PARAM_DXDESC
	DxdataPlaceholder            FusedOpsConstParamLabel = C.CUDNN_PARAM_DXDATA_PLACEHOLDER
	Dzdesc                       FusedOpsConstParamLabel = C.CUDNN_PARAM_DZDESC
	DzdataPlaceholder            FusedOpsConstParamLabel = C.CUDNN_PARAM_DZDATA_PLACEHOLDER
	BnDscalePlaceholder          FusedOpsConstParamLabel = C.CUDNN_PARAM_BN_DSCALE_PLACEHOLDER
	BnDbiasPlaceholder           FusedOpsConstParamLabel = C.CUDNN_PARAM_BN_DBIAS_PLACEHOLDER
)

// C returns the C representation of FusedOpsConstParamLabel
func (e FusedOpsConstParamLabel) C() C.cudnnFusedOpsConstParamLabel_t {
	return C.cudnnFusedOpsConstParamLabel_t(e)
}

type FusedOpsPointerPlaceHolder int

const (
	NullPtr        FusedOpsPointerPlaceHolder = C.CUDNN_PTR_NULL
	PtrElemAligned FusedOpsPointerPlaceHolder = C.CUDNN_PTR_ELEM_ALIGNED
	Ptr16          FusedOpsPointerPlaceHolder = C.CUDNN_PTR_16B_ALIGNED
)

// C returns the C representation of FusedOpsPointerPlaceHolder
func (e FusedOpsPointerPlaceHolder) C() C.cudnnFusedOpsPointerPlaceHolder_t {
	return C.cudnnFusedOpsPointerPlaceHolder_t(e)
}

type FusedOpsVariantParamLabel int

const (
	PtrXdata                        FusedOpsVariantParamLabel = C.CUDNN_PTR_XDATA
	PtrBnEqscale                    FusedOpsVariantParamLabel = C.CUDNN_PTR_BN_EQSCALE
	PtrBnEqbias                     FusedOpsVariantParamLabel = C.CUDNN_PTR_BN_EQBIAS
	PtrWdata                        FusedOpsVariantParamLabel = C.CUDNN_PTR_WDATA
	PtrDwdata                       FusedOpsVariantParamLabel = C.CUDNN_PTR_DWDATA
	PtrYdata                        FusedOpsVariantParamLabel = C.CUDNN_PTR_YDATA
	PtrDydata                       FusedOpsVariantParamLabel = C.CUDNN_PTR_DYDATA
	PtrYsum                         FusedOpsVariantParamLabel = C.CUDNN_PTR_YSUM
	PtrYsqsum                       FusedOpsVariantParamLabel = C.CUDNN_PTR_YSQSUM
	PtrWorkspace                    FusedOpsVariantParamLabel = C.CUDNN_PTR_WORKSPACE
	PtrBnScale                      FusedOpsVariantParamLabel = C.CUDNN_PTR_BN_SCALE
	PtrBnBias                       FusedOpsVariantParamLabel = C.CUDNN_PTR_BN_BIAS
	PtrBnSavedMean                  FusedOpsVariantParamLabel = C.CUDNN_PTR_BN_SAVED_MEAN
	PtrBnSavedInvstd                FusedOpsVariantParamLabel = C.CUDNN_PTR_BN_SAVED_INVSTD
	PtrBnRunningMean                FusedOpsVariantParamLabel = C.CUDNN_PTR_BN_RUNNING_MEAN
	PtrBnRunningVar                 FusedOpsVariantParamLabel = C.CUDNN_PTR_BN_RUNNING_VAR
	PtrZdata                        FusedOpsVariantParamLabel = C.CUDNN_PTR_ZDATA
	PtrBnZEqscale                   FusedOpsVariantParamLabel = C.CUDNN_PTR_BN_Z_EQSCALE
	PtrBnZEqbias                    FusedOpsVariantParamLabel = C.CUDNN_PTR_BN_Z_EQBIAS
	PtrActivationBitmask            FusedOpsVariantParamLabel = C.CUDNN_PTR_ACTIVATION_BITMASK
	PtrDxdata                       FusedOpsVariantParamLabel = C.CUDNN_PTR_DXDATA
	PtrDzdata                       FusedOpsVariantParamLabel = C.CUDNN_PTR_DZDATA
	PtrBnDscale                     FusedOpsVariantParamLabel = C.CUDNN_PTR_BN_DSCALE
	PtrBnDbias                      FusedOpsVariantParamLabel = C.CUDNN_PTR_BN_DBIAS
	ScalarSizeTWorkspaceSizeInBytes FusedOpsVariantParamLabel = C.CUDNN_SCALAR_SIZE_T_WORKSPACE_SIZE_IN_BYTES
	ScalarInt64TBnAccumulationCount FusedOpsVariantParamLabel = C.CUDNN_SCALAR_INT64_T_BN_ACCUMULATION_COUNT
	ScalarDoubleBnExpAvgFactor      FusedOpsVariantParamLabel = C.CUDNN_SCALAR_DOUBLE_BN_EXP_AVG_FACTOR
	ScalarDoubleBnEpsilon           FusedOpsVariantParamLabel = C.CUDNN_SCALAR_DOUBLE_BN_EPSILON
)

// C returns the C representation of FusedOpsVariantParamLabel
func (e FusedOpsVariantParamLabel) C() C.cudnnFusedOpsVariantParamLabel_t {
	return C.cudnnFusedOpsVariantParamLabel_t(e)
}

type FusedOps int

const (
	ScaleBiasActivationConvBnstats   FusedOps = C.CUDNN_FUSED_SCALE_BIAS_ACTIVATION_CONV_BNSTATS
	ScaleBiasActivationWgrad         FusedOps = C.CUDNN_FUSED_SCALE_BIAS_ACTIVATION_WGRAD
	BnFinalizeStatisticsTraining     FusedOps = C.CUDNN_FUSED_BN_FINALIZE_STATISTICS_TRAINING
	BnFinalizeStatisticsInference    FusedOps = C.CUDNN_FUSED_BN_FINALIZE_STATISTICS_INFERENCE
	ConvScaleBiasAddActivation       FusedOps = C.CUDNN_FUSED_CONV_SCALE_BIAS_ADD_ACTIVATION
	ScaleBiasAddActivationGenBitmask FusedOps = C.CUDNN_FUSED_SCALE_BIAS_ADD_ACTIVATION_GEN_BITMASK
	DactivationForkDbatchnorm        FusedOps = C.CUDNN_FUSED_DACTIVATION_FORK_DBATCHNORM
)

// C returns the C representation of FusedOps
func (e FusedOps) C() C.cudnnFusedOps_t { return C.cudnnFusedOps_t(e) }

type GenStatsMode int

const (
	SumSq GenStatsMode = C.CUDNN_GENSTATS_SUM_SQSUM
)

// C returns the C representation of GenStatsMode
func (e GenStatsMode) C() C.cudnnGenStatsMode_t { return C.cudnnGenStatsMode_t(e) }

type IndicesType int

const (
	Indices32 IndicesType = C.CUDNN_32BIT_INDICES
	Indices64 IndicesType = C.CUDNN_64BIT_INDICES
	Indices16 IndicesType = C.CUDNN_16BIT_INDICES
	Indices8  IndicesType = C.CUDNN_8BIT_INDICES
)

// C returns the C representation of IndicesType
func (e IndicesType) C() C.cudnnIndicesType_t { return C.cudnnIndicesType_t(e) }

type LRNMode int

const (
	CrossChannelDim1 LRNMode = C.CUDNN_LRN_CROSS_CHANNEL_DIM1
)

// C returns the C representation of LRNMode
func (e LRNMode) C() C.cudnnLRNMode_t { return C.cudnnLRNMode_t(e) }

type LossNormalizationMode int

const (
	None    LossNormalizationMode = C.CUDNN_LOSS_NORMALIZATION_NONE
	Softmax LossNormalizationMode = C.CUDNN_LOSS_NORMALIZATION_SOFTMAX
)

// C returns the C representation of LossNormalizationMode
func (e LossNormalizationMode) C() C.cudnnLossNormalizationMode_t {
	return C.cudnnLossNormalizationMode_t(e)
}

type MathType int

const (
	DefaultMath                 MathType = C.CUDNN_DEFAULT_MATH
	TensorOpMath                MathType = C.CUDNN_TENSOR_OP_MATH
	TensorOpMathAllowConversion MathType = C.CUDNN_TENSOR_OP_MATH_ALLOW_CONVERSION
	FmaMath                     MathType = C.CUDNN_FMA_MATH
)

// C returns the C representation of MathType
func (e MathType) C() C.cudnnMathType_t { return C.cudnnMathType_t(e) }

type MultiHeadAttnWeightKind int

const (
	QWeights MultiHeadAttnWeightKind = C.CUDNN_MH_ATTN_Q_WEIGHTS
	KWeights MultiHeadAttnWeightKind = C.CUDNN_MH_ATTN_K_WEIGHTS
	VWeights MultiHeadAttnWeightKind = C.CUDNN_MH_ATTN_V_WEIGHTS
	OWeights MultiHeadAttnWeightKind = C.CUDNN_MH_ATTN_O_WEIGHTS
	QBiases  MultiHeadAttnWeightKind = C.CUDNN_MH_ATTN_Q_BIASES
	KBiases  MultiHeadAttnWeightKind = C.CUDNN_MH_ATTN_K_BIASES
	VBiases  MultiHeadAttnWeightKind = C.CUDNN_MH_ATTN_V_BIASES
	OBiases  MultiHeadAttnWeightKind = C.CUDNN_MH_ATTN_O_BIASES
)

// C returns the C representation of MultiHeadAttnWeightKind
func (e MultiHeadAttnWeightKind) C() C.cudnnMultiHeadAttnWeightKind_t {
	return C.cudnnMultiHeadAttnWeightKind_t(e)
}

type NanPropagation int

const (
	NotPropagateNan NanPropagation = C.CUDNN_NOT_PROPAGATE_NAN
	PropagateNan    NanPropagation = C.CUDNN_PROPAGATE_NAN
)

// C returns the C representation of NanPropagation
func (e NanPropagation) C() C.cudnnNanPropagation_t { return C.cudnnNanPropagation_t(e) }

type NormAlgo int

const (
	Standard NormAlgo = C.CUDNN_NORM_ALGO_STANDARD
	Persist  NormAlgo = C.CUDNN_NORM_ALGO_PERSIST
)

// C returns the C representation of NormAlgo
func (e NormAlgo) C() C.cudnnNormAlgo_t { return C.cudnnNormAlgo_t(e) }

type NormMode int

const (
	Activation NormMode = C.CUDNN_NORM_PER_ACTIVATION
	Channel    NormMode = C.CUDNN_NORM_PER_CHANNEL
)

// C returns the C representation of NormMode
func (e NormMode) C() C.cudnnNormMode_t { return C.cudnnNormMode_t(e) }

type NormOps int

const (
	Norm          NormOps = C.CUDNN_NORM_OPS_NORM
	Activation    NormOps = C.CUDNN_NORM_OPS_NORM_ACTIVATION
	AddActivation NormOps = C.CUDNN_NORM_OPS_NORM_ADD_ACTIVATION
)

// C returns the C representation of NormOps
func (e NormOps) C() C.cudnnNormOps_t { return C.cudnnNormOps_t(e) }

type OpTensorOp int

const (
	Add  OpTensorOp = C.CUDNN_OP_TENSOR_ADD
	Mul  OpTensorOp = C.CUDNN_OP_TENSOR_MUL
	Min  OpTensorOp = C.CUDNN_OP_TENSOR_MIN
	Max  OpTensorOp = C.CUDNN_OP_TENSOR_MAX
	Sqrt OpTensorOp = C.CUDNN_OP_TENSOR_SQRT
	Not  OpTensorOp = C.CUDNN_OP_TENSOR_NOT
)

// C returns the C representation of OpTensorOp
func (e OpTensorOp) C() C.cudnnOpTensorOp_t { return C.cudnnOpTensorOp_t(e) }

type PointwiseMode int

const (
	Add        PointwiseMode = C.CUDNN_POINTWISE_ADD
	Mul        PointwiseMode = C.CUDNN_POINTWISE_MUL
	Min        PointwiseMode = C.CUDNN_POINTWISE_MIN
	Max        PointwiseMode = C.CUDNN_POINTWISE_MAX
	Sqrt       PointwiseMode = C.CUDNN_POINTWISE_SQRT
	ReluFwd    PointwiseMode = C.CUDNN_POINTWISE_RELU_FWD
	TanhFwd    PointwiseMode = C.CUDNN_POINTWISE_TANH_FWD
	SigmoidFwd PointwiseMode = C.CUDNN_POINTWISE_SIGMOID_FWD
	EluFwd     PointwiseMode = C.CUDNN_POINTWISE_ELU_FWD
)

// C returns the C representation of PointwiseMode
func (e PointwiseMode) C() C.cudnnPointwiseMode_t { return C.cudnnPointwiseMode_t(e) }

type PoolingMode int

const (
	MaxPooling                 PoolingMode = C.CUDNN_POOLING_MAX
	AverageCountIncludePadding PoolingMode = C.CUDNN_POOLING_AVERAGE_COUNT_INCLUDE_PADDING
	AverageCountExcludePadding PoolingMode = C.CUDNN_POOLING_AVERAGE_COUNT_EXCLUDE_PADDING
	MaxDeterministic           PoolingMode = C.CUDNN_POOLING_MAX_DETERMINISTIC
)

// C returns the C representation of PoolingMode
func (e PoolingMode) C() C.cudnnPoolingMode_t { return C.cudnnPoolingMode_t(e) }

type RNNAlgo int

const (
	Standard       RNNAlgo = C.CUDNN_RNN_ALGO_STANDARD
	PersistStatic  RNNAlgo = C.CUDNN_RNN_ALGO_PERSIST_STATIC
	PersistDynamic RNNAlgo = C.CUDNN_RNN_ALGO_PERSIST_DYNAMIC
	Count          RNNAlgo = C.CUDNN_RNN_ALGO_COUNT
)

// C returns the C representation of RNNAlgo
func (e RNNAlgo) C() C.cudnnRNNAlgo_t { return C.cudnnRNNAlgo_t(e) }

type RNNBiasMode int

const (
	NoBias        RNNBiasMode = C.CUDNN_RNN_NO_BIAS
	SingleInpBias RNNBiasMode = C.CUDNN_RNN_SINGLE_INP_BIAS
	DoubleBias    RNNBiasMode = C.CUDNN_RNN_DOUBLE_BIAS
	SingleRecBias RNNBiasMode = C.CUDNN_RNN_SINGLE_REC_BIAS
)

// C returns the C representation of RNNBiasMode
func (e RNNBiasMode) C() C.cudnnRNNBiasMode_t { return C.cudnnRNNBiasMode_t(e) }

type RNNClipMode int

const (
	None   RNNClipMode = C.CUDNN_RNN_CLIP_NONE
	Minmax RNNClipMode = C.CUDNN_RNN_CLIP_MINMAX
)

// C returns the C representation of RNNClipMode
func (e RNNClipMode) C() C.cudnnRNNClipMode_t { return C.cudnnRNNClipMode_t(e) }

type RNNDataLayout int

const (
	SeqMajorUnpacked   RNNDataLayout = C.CUDNN_RNN_DATA_LAYOUT_SEQ_MAJOR_UNPACKED
	SeqMajorPacked     RNNDataLayout = C.CUDNN_RNN_DATA_LAYOUT_SEQ_MAJOR_PACKED
	BatchMajorUnpacked RNNDataLayout = C.CUDNN_RNN_DATA_LAYOUT_BATCH_MAJOR_UNPACKED
)

// C returns the C representation of RNNDataLayout
func (e RNNDataLayout) C() C.cudnnRNNDataLayout_t { return C.cudnnRNNDataLayout_t(e) }

type RNNInputMode int

const (
	LinearInput RNNInputMode = C.CUDNN_LINEAR_INPUT
	SkipInput   RNNInputMode = C.CUDNN_SKIP_INPUT
)

// C returns the C representation of RNNInputMode
func (e RNNInputMode) C() C.cudnnRNNInputMode_t { return C.cudnnRNNInputMode_t(e) }

type RNNMode int

const (
	RNNReLU RNNMode = C.CUDNN_RNN_RELU
	RNNTanh RNNMode = C.CUDNN_RNN_TANH
	LSTM    RNNMode = C.CUDNN_LSTM
	GRU     RNNMode = C.CUDNN_GRU
)

// C returns the C representation of RNNMode
func (e RNNMode) C() C.cudnnRNNMode_t { return C.cudnnRNNMode_t(e) }

type ReduceTensorIndices int

const (
	ReduceNoIndices        ReduceTensorIndices = C.CUDNN_REDUCE_TENSOR_NO_INDICES
	ReduceFlattenedIndices ReduceTensorIndices = C.CUDNN_REDUCE_TENSOR_FLATTENED_INDICES
)

// C returns the C representation of ReduceTensorIndices
func (e ReduceTensorIndices) C() C.cudnnReduceTensorIndices_t { return C.cudnnReduceTensorIndices_t(e) }

type ReduceTensorOp int

const (
	ReduceAdd        ReduceTensorOp = C.CUDNN_REDUCE_TENSOR_ADD
	ReduceMul        ReduceTensorOp = C.CUDNN_REDUCE_TENSOR_MUL
	ReduceMin        ReduceTensorOp = C.CUDNN_REDUCE_TENSOR_MIN
	ReduceMax        ReduceTensorOp = C.CUDNN_REDUCE_TENSOR_MAX
	ReduceAmax       ReduceTensorOp = C.CUDNN_REDUCE_TENSOR_AMAX
	ReduceAvg        ReduceTensorOp = C.CUDNN_REDUCE_TENSOR_AVG
	ReduceNorm1      ReduceTensorOp = C.CUDNN_REDUCE_TENSOR_NORM1
	ReduceNorm2      ReduceTensorOp = C.CUDNN_REDUCE_TENSOR_NORM2
	ReduceMulNoZeros ReduceTensorOp = C.CUDNN_REDUCE_TENSOR_MUL_NO_ZEROS
)

// C returns the C representation of ReduceTensorOp
func (e ReduceTensorOp) C() C.cudnnReduceTensorOp_t { return C.cudnnReduceTensorOp_t(e) }

type ReorderType int

const (
	DefaultReorder ReorderType = C.CUDNN_DEFAULT_REORDER
	NoReorder      ReorderType = C.CUDNN_NO_REORDER
)

// C returns the C representation of ReorderType
func (e ReorderType) C() C.cudnnReorderType_t { return C.cudnnReorderType_t(e) }

type SamplerType int

const (
	Bilinear SamplerType = C.CUDNN_SAMPLER_BILINEAR
)

// C returns the C representation of SamplerType
func (e SamplerType) C() C.cudnnSamplerType_t { return C.cudnnSamplerType_t(e) }

type SeqDataAxis int

const (
	TimeDim  SeqDataAxis = C.CUDNN_SEQDATA_TIME_DIM
	BatchDim SeqDataAxis = C.CUDNN_SEQDATA_BATCH_DIM
	BeamDim  SeqDataAxis = C.CUDNN_SEQDATA_BEAM_DIM
	VectDim  SeqDataAxis = C.CUDNN_SEQDATA_VECT_DIM
)

// C returns the C representation of SeqDataAxis
func (e SeqDataAxis) C() C.cudnnSeqDataAxis_t { return C.cudnnSeqDataAxis_t(e) }

type Severity int

const (
	Fatal   Severity = C.CUDNN_SEV_FATAL
	Error   Severity = C.CUDNN_SEV_ERROR
	Warning Severity = C.CUDNN_SEV_WARNING
	Info    Severity = C.CUDNN_SEV_INFO
)

// C returns the C representation of Severity
func (e Severity) C() C.cudnnSeverity_t { return C.cudnnSeverity_t(e) }

type SoftmaxAlgorithm int

const (
	Fast     SoftmaxAlgorithm = C.CUDNN_SOFTMAX_FAST
	Accurate SoftmaxAlgorithm = C.CUDNN_SOFTMAX_ACCURATE
	Log      SoftmaxAlgorithm = C.CUDNN_SOFTMAX_LOG
)

// C returns the C representation of SoftmaxAlgorithm
func (e SoftmaxAlgorithm) C() C.cudnnSoftmaxAlgorithm_t { return C.cudnnSoftmaxAlgorithm_t(e) }

type SoftmaxMode int

const (
	Instance SoftmaxMode = C.CUDNN_SOFTMAX_MODE_INSTANCE
	Channel  SoftmaxMode = C.CUDNN_SOFTMAX_MODE_CHANNEL
)

// C returns the C representation of SoftmaxMode
func (e SoftmaxMode) C() C.cudnnSoftmaxMode_t { return C.cudnnSoftmaxMode_t(e) }

type TensorFormat int

const (
	NCHW      TensorFormat = C.CUDNN_TENSOR_NCHW
	NHWC      TensorFormat = C.CUDNN_TENSOR_NHWC
	NCHWVectC TensorFormat = C.CUDNN_TENSOR_NCHW_VECT_C
)

// C returns the C representation of TensorFormat
func (e TensorFormat) C() C.cudnnTensorFormat_t { return C.cudnnTensorFormat_t(e) }

type WgradMode int

const (
	Add WgradMode = C.CUDNN_WGRAD_MODE_ADD
	Set WgradMode = C.CUDNN_WGRAD_MODE_SET
)

// C returns the C representation of WgradMode
func (e WgradMode) C() C.cudnnWgradMode_t { return C.cudnnWgradMode_t(e) }
