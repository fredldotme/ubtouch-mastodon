/*
 * Copyright (C) 2022  Development@bendingtherules.nl
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; version 3.
 *
 * mastodon is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
	"mastodon-client/mastodon"
	"mastodon-client/pushnotifications"
	"os"
)

type RootObject struct {
	core.QObject
	_ *pushnotifications.PushHandler `property:"pushNotifications"`
}

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var view = quick.NewQQuickView(nil)

	rootObject := NewRootObject(nil)
	rootObject.SetPushNotifications(pushnotifications.NewPushHandler(nil))
	qClient, _ := mastodon.GetQClient()

	view.RootContext().SetContextProperty("rootObject", rootObject)
	view.RootContext().SetContextProperty("QClient", qClient)

	view.SetSource(core.NewQUrl3("qrc:/qml/Main.qml", 0))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.Show()

	gui.QGuiApplication_Exec()
}
