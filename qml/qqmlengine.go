package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"strings"
	"unsafe"
)

type QQmlEngine struct {
	QJSEngine
}

type QQmlEngine_ITF interface {
	QJSEngine_ITF
	QQmlEngine_PTR() *QQmlEngine
}

func PointerFromQQmlEngine(ptr QQmlEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlEngine_PTR().Pointer()
	}
	return nil
}

func NewQQmlEngineFromPointer(ptr unsafe.Pointer) *QQmlEngine {
	var n = new(QQmlEngine)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QQmlEngine_") {
		n.SetObjectName("QQmlEngine_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlEngine) QQmlEngine_PTR() *QQmlEngine {
	return ptr
}

//QQmlEngine::ObjectOwnership
type QQmlEngine__ObjectOwnership int64

const (
	QQmlEngine__CppOwnership        = QQmlEngine__ObjectOwnership(0)
	QQmlEngine__JavaScriptOwnership = QQmlEngine__ObjectOwnership(1)
)

func (ptr *QQmlEngine) OfflineStoragePath() string {
	defer qt.Recovering("QQmlEngine::offlineStoragePath")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlEngine_OfflineStoragePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlEngine) SetOfflineStoragePath(dir string) {
	defer qt.Recovering("QQmlEngine::setOfflineStoragePath")

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetOfflineStoragePath(ptr.Pointer(), C.CString(dir))
	}
}

func NewQQmlEngine(parent core.QObject_ITF) *QQmlEngine {
	defer qt.Recovering("QQmlEngine::QQmlEngine")

	return NewQQmlEngineFromPointer(C.QQmlEngine_NewQQmlEngine(core.PointerFromQObject(parent)))
}

func (ptr *QQmlEngine) AddImageProvider(providerId string, provider QQmlImageProviderBase_ITF) {
	defer qt.Recovering("QQmlEngine::addImageProvider")

	if ptr.Pointer() != nil {
		C.QQmlEngine_AddImageProvider(ptr.Pointer(), C.CString(providerId), PointerFromQQmlImageProviderBase(provider))
	}
}

func (ptr *QQmlEngine) AddImportPath(path string) {
	defer qt.Recovering("QQmlEngine::addImportPath")

	if ptr.Pointer() != nil {
		C.QQmlEngine_AddImportPath(ptr.Pointer(), C.CString(path))
	}
}

func (ptr *QQmlEngine) AddPluginPath(path string) {
	defer qt.Recovering("QQmlEngine::addPluginPath")

	if ptr.Pointer() != nil {
		C.QQmlEngine_AddPluginPath(ptr.Pointer(), C.CString(path))
	}
}

func (ptr *QQmlEngine) ClearComponentCache() {
	defer qt.Recovering("QQmlEngine::clearComponentCache")

	if ptr.Pointer() != nil {
		C.QQmlEngine_ClearComponentCache(ptr.Pointer())
	}
}

func QQmlEngine_ContextForObject(object core.QObject_ITF) *QQmlContext {
	defer qt.Recovering("QQmlEngine::contextForObject")

	return NewQQmlContextFromPointer(C.QQmlEngine_QQmlEngine_ContextForObject(core.PointerFromQObject(object)))
}

func (ptr *QQmlEngine) ImageProvider(providerId string) *QQmlImageProviderBase {
	defer qt.Recovering("QQmlEngine::imageProvider")

	if ptr.Pointer() != nil {
		return NewQQmlImageProviderBaseFromPointer(C.QQmlEngine_ImageProvider(ptr.Pointer(), C.CString(providerId)))
	}
	return nil
}

func (ptr *QQmlEngine) ImportPathList() []string {
	defer qt.Recovering("QQmlEngine::importPathList")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QQmlEngine_ImportPathList(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QQmlEngine) IncubationController() *QQmlIncubationController {
	defer qt.Recovering("QQmlEngine::incubationController")

	if ptr.Pointer() != nil {
		return NewQQmlIncubationControllerFromPointer(C.QQmlEngine_IncubationController(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) NetworkAccessManager() *network.QNetworkAccessManager {
	defer qt.Recovering("QQmlEngine::networkAccessManager")

	if ptr.Pointer() != nil {
		return network.NewQNetworkAccessManagerFromPointer(C.QQmlEngine_NetworkAccessManager(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) NetworkAccessManagerFactory() *QQmlNetworkAccessManagerFactory {
	defer qt.Recovering("QQmlEngine::networkAccessManagerFactory")

	if ptr.Pointer() != nil {
		return NewQQmlNetworkAccessManagerFactoryFromPointer(C.QQmlEngine_NetworkAccessManagerFactory(ptr.Pointer()))
	}
	return nil
}

func QQmlEngine_ObjectOwnership(object core.QObject_ITF) QQmlEngine__ObjectOwnership {
	defer qt.Recovering("QQmlEngine::objectOwnership")

	return QQmlEngine__ObjectOwnership(C.QQmlEngine_QQmlEngine_ObjectOwnership(core.PointerFromQObject(object)))
}

func (ptr *QQmlEngine) OutputWarningsToStandardError() bool {
	defer qt.Recovering("QQmlEngine::outputWarningsToStandardError")

	if ptr.Pointer() != nil {
		return C.QQmlEngine_OutputWarningsToStandardError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlEngine) PluginPathList() []string {
	defer qt.Recovering("QQmlEngine::pluginPathList")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QQmlEngine_PluginPathList(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QQmlEngine) ConnectQuit(f func()) {
	defer qt.Recovering("connect QQmlEngine::quit")

	if ptr.Pointer() != nil {
		C.QQmlEngine_ConnectQuit(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "quit", f)
	}
}

func (ptr *QQmlEngine) DisconnectQuit() {
	defer qt.Recovering("disconnect QQmlEngine::quit")

	if ptr.Pointer() != nil {
		C.QQmlEngine_DisconnectQuit(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "quit")
	}
}

//export callbackQQmlEngineQuit
func callbackQQmlEngineQuit(ptrName *C.char) {
	defer qt.Recovering("callback QQmlEngine::quit")

	var signal = qt.GetSignal(C.GoString(ptrName), "quit")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlEngine) RemoveImageProvider(providerId string) {
	defer qt.Recovering("QQmlEngine::removeImageProvider")

	if ptr.Pointer() != nil {
		C.QQmlEngine_RemoveImageProvider(ptr.Pointer(), C.CString(providerId))
	}
}

func (ptr *QQmlEngine) RootContext() *QQmlContext {
	defer qt.Recovering("QQmlEngine::rootContext")

	if ptr.Pointer() != nil {
		return NewQQmlContextFromPointer(C.QQmlEngine_RootContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) SetBaseUrl(url core.QUrl_ITF) {
	defer qt.Recovering("QQmlEngine::setBaseUrl")

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetBaseUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func QQmlEngine_SetContextForObject(object core.QObject_ITF, context QQmlContext_ITF) {
	defer qt.Recovering("QQmlEngine::setContextForObject")

	C.QQmlEngine_QQmlEngine_SetContextForObject(core.PointerFromQObject(object), PointerFromQQmlContext(context))
}

func (ptr *QQmlEngine) SetImportPathList(paths []string) {
	defer qt.Recovering("QQmlEngine::setImportPathList")

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetImportPathList(ptr.Pointer(), C.CString(strings.Join(paths, ",,,")))
	}
}

func (ptr *QQmlEngine) SetIncubationController(controller QQmlIncubationController_ITF) {
	defer qt.Recovering("QQmlEngine::setIncubationController")

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetIncubationController(ptr.Pointer(), PointerFromQQmlIncubationController(controller))
	}
}

func (ptr *QQmlEngine) SetNetworkAccessManagerFactory(factory QQmlNetworkAccessManagerFactory_ITF) {
	defer qt.Recovering("QQmlEngine::setNetworkAccessManagerFactory")

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetNetworkAccessManagerFactory(ptr.Pointer(), PointerFromQQmlNetworkAccessManagerFactory(factory))
	}
}

func QQmlEngine_SetObjectOwnership(object core.QObject_ITF, ownership QQmlEngine__ObjectOwnership) {
	defer qt.Recovering("QQmlEngine::setObjectOwnership")

	C.QQmlEngine_QQmlEngine_SetObjectOwnership(core.PointerFromQObject(object), C.int(ownership))
}

func (ptr *QQmlEngine) SetOutputWarningsToStandardError(enabled bool) {
	defer qt.Recovering("QQmlEngine::setOutputWarningsToStandardError")

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetOutputWarningsToStandardError(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QQmlEngine) SetPluginPathList(paths []string) {
	defer qt.Recovering("QQmlEngine::setPluginPathList")

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetPluginPathList(ptr.Pointer(), C.CString(strings.Join(paths, ",,,")))
	}
}

func (ptr *QQmlEngine) TrimComponentCache() {
	defer qt.Recovering("QQmlEngine::trimComponentCache")

	if ptr.Pointer() != nil {
		C.QQmlEngine_TrimComponentCache(ptr.Pointer())
	}
}

func (ptr *QQmlEngine) DestroyQQmlEngine() {
	defer qt.Recovering("QQmlEngine::~QQmlEngine")

	if ptr.Pointer() != nil {
		C.QQmlEngine_DestroyQQmlEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}