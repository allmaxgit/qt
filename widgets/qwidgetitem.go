package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QWidgetItem struct {
	QLayoutItem
}

type QWidgetItem_ITF interface {
	QLayoutItem_ITF
	QWidgetItem_PTR() *QWidgetItem
}

func PointerFromQWidgetItem(ptr QWidgetItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidgetItem_PTR().Pointer()
	}
	return nil
}

func NewQWidgetItemFromPointer(ptr unsafe.Pointer) *QWidgetItem {
	var n = new(QWidgetItem)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QWidgetItem_") {
		n.SetObjectNameAbs("QWidgetItem_" + qt.Identifier())
	}
	return n
}

func (ptr *QWidgetItem) QWidgetItem_PTR() *QWidgetItem {
	return ptr
}

func NewQWidgetItem(widget QWidget_ITF) *QWidgetItem {
	defer qt.Recovering("QWidgetItem::QWidgetItem")

	return NewQWidgetItemFromPointer(C.QWidgetItem_NewQWidgetItem(PointerFromQWidget(widget)))
}

func (ptr *QWidgetItem) ControlTypes() QSizePolicy__ControlType {
	defer qt.Recovering("QWidgetItem::controlTypes")

	if ptr.Pointer() != nil {
		return QSizePolicy__ControlType(C.QWidgetItem_ControlTypes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidgetItem) ExpandingDirections() core.Qt__Orientation {
	defer qt.Recovering("QWidgetItem::expandingDirections")

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QWidgetItem_ExpandingDirections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidgetItem) HasHeightForWidth() bool {
	defer qt.Recovering("QWidgetItem::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QWidgetItem_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidgetItem) HeightForWidth(w int) int {
	defer qt.Recovering("QWidgetItem::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QWidgetItem_HeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QWidgetItem) IsEmpty() bool {
	defer qt.Recovering("QWidgetItem::isEmpty")

	if ptr.Pointer() != nil {
		return C.QWidgetItem_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidgetItem) Widget() *QWidget {
	defer qt.Recovering("QWidgetItem::widget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidgetItem_Widget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidgetItem) DestroyQWidgetItem() {
	defer qt.Recovering("QWidgetItem::~QWidgetItem")

	if ptr.Pointer() != nil {
		C.QWidgetItem_DestroyQWidgetItem(ptr.Pointer())
	}
}

func (ptr *QWidgetItem) ObjectNameAbs() string {
	defer qt.Recovering("QWidgetItem::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWidgetItem_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidgetItem) SetObjectNameAbs(name string) {
	defer qt.Recovering("QWidgetItem::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QWidgetItem_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}