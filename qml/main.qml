import QtQuick 2.0
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.1
import QtQuick.Layouts 1.3

ApplicationWindow {
    visible: true
    title: qsTr("Radix converter")

    Material.theme: Material.Dark
    Material.accent: Material.Teal

    minimumWidth: 820
    minimumHeight: 480
    maximumWidth: minimumWidth
    maximumHeight: minimumHeight

    ColumnLayout {
        anchors.fill: parent

        RowLayout {
            Layout.fillWidth: true

            Label {
                text: qsTr("Number to convert")
                font.pointSize: 12
                padding: 0
                Layout.fillWidth: false
                leftPadding: 10
                horizontalAlignment: Text.AlignLeft
                verticalAlignment: Text.AlignVCenter
            }

            TextField {
                id: targetNumberField
                horizontalAlignment: Text.AlignHCenter
                font.pointSize: 12
                Layout.fillWidth: true
                placeholderText: qsTr("Enter here your number")
                maximumLength: 15
                focus: true
                validator: RegExpValidator {
                    regExp: /[0-9a-fA-F]+/
                }

                onTextChanged: {
                    if (targetNumberField.length >= 1) {
                        convertButton.enabled = true;
                    } else {
                        convertButton.enabled = false;
                    }
                }
            }
        }

        RowLayout {
            ColumnLayout {
                Label {
                    text: qsTr("From")
                    font.pointSize: 12
                    horizontalAlignment: Text.AlignHCenter
                    Layout.fillWidth: true
                }

                ComboBox {
                    id: fromRadixBox
                    Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter

                    Component.onCompleted: {
                        var vals = []
                        for (var i = 2; i <= 16; i++) {
                            vals.push(i);
                        }
                        model = vals
                    }
                }
            }

            Button {
                id: switchBases
                text: "Switch"

                onClicked: {
                    var toRadixBoxIndex = toRadixBox.currentIndex;
                    toRadixBox.currentIndex = fromRadixBox.currentIndex;
                    fromRadixBox.currentIndex = toRadixBoxIndex;
                }
            }

            ColumnLayout {
                Label {
                    text: qsTr("To")
                    font.pointSize: 12
                    horizontalAlignment: Text.AlignHCenter
                    Layout.fillWidth: true
                }

                ComboBox {
                    id: toRadixBox
                    model: []
                    Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter

                    Component.onCompleted: {
                        var vals = []
                        for (var i = 2; i <= 16; i++) {
                            vals.push(i);
                        }
                        model = vals
                        currentIndex = 14;
                    }
                }
            }

        }

        RowLayout {
            Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
            Layout.fillWidth: true

            Button {
                id: convertButton
                enabled: true
                text: qsTr("Convert")
                Layout.fillWidth: false
                Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter

                onClicked: {
                    resultField.text = _converter.convertButtonClicked(targetNumberField.text, parseInt(fromRadixBox.currentText), parseInt(toRadixBox.currentText));
                }
            }
        }

        RowLayout {
            Label {
                text: qsTr("Result")
                font.pointSize: 12
                horizontalAlignment: Text.AlignLeft
                leftPadding: 10
            }

            TextField {
                id: resultField
                text: qsTr("")
                Layout.fillWidth: true
                Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
                placeholderText: "Here you'll see errors or result mesage"
                horizontalAlignment: Text.AlignHCenter
                readOnly: true
            }
        }
    }
}