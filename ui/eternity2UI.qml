import QtQuick 2.1
import QtQuick.Controls 1.0
import QtQuick.Window 2.0
import QtQuick.Dialogs 1.0

ApplicationWindow {
    id: appWin
    width: 800
    height: 600
    title: qsTr("Image Grid")

    Grid {
        id: grid1
		objectName: "grid"
        x: 8
        y: 80
        width: 512
        height: 512
        rows: 16
        columns: 16

        Image {
            id: image_1_1
            objectName: "tile_1_1"
            x: 69
            y: 182
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_1_2
            objectName: "tile_1_2"
            x: 76
            y: 185
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_1_3
            objectName: "tile_1_3"
            x: 74
            y: 173
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_1_4
            objectName: "tile_1_4"
            x: 75
            y: 181
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_1_5
            objectName: "tile_1_5"
            x: 74
            y: 184
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_1_6
            objectName: "tile_1_6"
            x: 61
            y: 186
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_1_7
            objectName: "tile_1_7"
            x: 67
            y: 178
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_1_8
            objectName: "tile_1_8"
            x: 60
            y: 191
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_1_9
            objectName: "tile_1_9"
            x: 75
            y: 183
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_1_10
            objectName: "tile_1_10"
            x: 73
            y: 186
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_1_11
            objectName: "tile_1_11"
            x: 78
            y: 183
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_1_12
            objectName: "tile_1_12"
            x: 77
            y: 173
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_1_13
            objectName: "tile_1_13"
            x: 67
            y: 175
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_1_14
            objectName: "tile_1_14"
            x: 60
            y: 180
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_1_15
            objectName: "tile_1_15"
            x: 74
            y: 191
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_1_16
            objectName: "tile_1_16"
            x: 59
            y: 180
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_2_1
            objectName: "tile_2_1"
            x: 74
            y: 184
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_2_2
            objectName: "tile_2_2"
            x: 76
            y: 185
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_2_3
            objectName: "tile_2_3"
            x: 78
            y: 175
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_2_4
            objectName: "tile_2_4"
            x: 68
            y: 176
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_2_5
            objectName: "tile_2_5"
            x: 59
            y: 174
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_2_6
            objectName: "tile_2_6"
            x: 64
            y: 182
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_2_7
            objectName: "tile_2_7"
            x: 68
            y: 188
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_2_8
            objectName: "tile_2_8"
            x: 78
            y: 175
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_2_9
            objectName: "tile_2_9"
            x: 69
            y: 184
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_2_10
            objectName: "tile_2_10"
            x: 76
            y: 188
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_2_11
            objectName: "tile_2_11"
            x: 63
            y: 181
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_2_12
            objectName: "tile_2_12"
            x: 61
            y: 173
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_2_13
            objectName: "tile_2_13"
            x: 76
            y: 175
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_2_14
            objectName: "tile_2_14"
            x: 69
            y: 172
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_2_15
            objectName: "tile_2_15"
            x: 66
            y: 179
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_2_16
            objectName: "tile_2_16"
            x: 59
            y: 172
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_3_1
            objectName: "tile_3_1"
            x: 66
            y: 190
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_3_2
            objectName: "tile_3_2"
            x: 65
            y: 188
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_3_3
            objectName: "tile_3_3"
            x: 71
            y: 186
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_3_4
            objectName: "tile_3_4"
            x: 68
            y: 177
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_3_5
            objectName: "tile_3_5"
            x: 69
            y: 189
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_3_6
            objectName: "tile_3_6"
            x: 73
            y: 174
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_3_7
            objectName: "tile_3_7"
            x: 78
            y: 185
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_3_8
            objectName: "tile_3_8"
            x: 78
            y: 186
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_3_9
            objectName: "tile_3_9"
            x: 66
            y: 191
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_3_10
            objectName: "tile_3_10"
            x: 72
            y: 172
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_3_11
            objectName: "tile_3_11"
            x: 60
            y: 191
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_3_12
            objectName: "tile_3_12"
            x: 66
            y: 173
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_3_13
            objectName: "tile_3_13"
            x: 69
            y: 174
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_3_14
            objectName: "tile_3_14"
            x: 63
            y: 174
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_3_15
            objectName: "tile_3_15"
            x: 65
            y: 190
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_3_16
            objectName: "tile_3_16"
            x: 76
            y: 180
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_4_1
            objectName: "tile_4_1"
            x: 73
            y: 181
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_4_2
            objectName: "tile_4_2"
            x: 74
            y: 184
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_4_3
            objectName: "tile_4_3"
            x: 63
            y: 179
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_4_4
            objectName: "tile_4_4"
            x: 76
            y: 186
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_4_5
            objectName: "tile_4_5"
            x: 78
            y: 189
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_4_6
            objectName: "tile_4_6"
            x: 74
            y: 187
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_4_7
            objectName: "tile_4_7"
            x: 63
            y: 175
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_4_8
            objectName: "tile_4_8"
            x: 69
            y: 184
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_4_9
            objectName: "tile_4_9"
            x: 70
            y: 179
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_4_10
            objectName: "tile_4_10"
            x: 68
            y: 184
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_4_11
            objectName: "tile_4_11"
            x: 78
            y: 174
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_4_12
            objectName: "tile_4_12"
            x: 71
            y: 175
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_4_13
            objectName: "tile_4_13"
            x: 64
            y: 184
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_4_14
            objectName: "tile_4_14"
            x: 75
            y: 177
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_4_15
            objectName: "tile_4_15"
            x: 65
            y: 186
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_4_16
            objectName: "tile_4_16"
            x: 67
            y: 175
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_5_1
			objectName: "tile_5_1"
            x: 71
            y: 172
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_5_2
			objectName: "tile_5_2"
            x: 65
            y: 173
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_5_3
            objectName: "tile_5_3"
			x: 61
            y: 183
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_5_4
			objectName: "tile_5_4"
            x: 72
            y: 185
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_5_5
			objectName: "tile_5_5"
            x: 69
            y: 176
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_5_6
			objectName: "tile_5_6"
            x: 65
            y: 177
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_5_7
			objectName: "tile_5_7"
            x: 65
            y: 175
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_5_8
			objectName: "tile_5_8"
            x: 59
            y: 186
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_5_9
			objectName: "tile_5_9"
            x: 71
            y: 181
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_5_10
			objectName: "tile_5_10"
            x: 60
            y: 176
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_5_11
			objectName: "tile_5_11"
            x: 60
            y: 183
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_5_12
			objectName: "tile_5_12"
            x: 59
            y: 180
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_5_13
			objectName: "tile_5_13"
            x: 65
            y: 179
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_5_14
			objectName: "tile_5_14"
            x: 65
            y: 178
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_5_15
			objectName: "tile_5_15"
            x: 77
            y: 182
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_5_16
			objectName: "tile_5_16"
            x: 72
            y: 183
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_6_1
			objectName: "tile_6_1"
            x: 63
            y: 180
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_6_2
			objectName: "tile_6_2"
            x: 70
            y: 183
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_6_3
			objectName: "tile_6_3"
            x: 66
            y: 179
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_6_4
			objectName: "tile_6_4"
            x: 60
            y: 173
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_6_5
			objectName: "tile_6_5"
            x: 66
            y: 177
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_6_6
			objectName: "tile_6_6"
            x: 72
            y: 186
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_6_7
			objectName: "tile_6_7"
            x: 63
            y: 178
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_6_8
			objectName: "tile_6_8"
            x: 78
            y: 189
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_6_9
			objectName: "tile_6_9"
            x: 64
            y: 178
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_6_10
			objectName: "tile_6_10"
            x: 77
            y: 172
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_6_11
			objectName: "tile_6_11"
            x: 65
            y: 187
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_6_12
			objectName: "tile_6_12"
            x: 71
            y: 185
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_6_13
			objectName: "tile_6_13"
            x: 60
            y: 176
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_6_14
			objectName: "tile_6_14"
            x: 68
            y: 185
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_6_15
			objectName: "tile_6_15"
            x: 66
            y: 187
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_6_16
			objectName: "tile_6_16"
            x: 71
            y: 189
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_7_1
			objectName: "tile_7_1"
            x: 76
            y: 185
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_7_2
			objectName: "tile_7_2"
            x: 72
            y: 190
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_7_3
			objectName: "tile_7_3"
            x: 74
            y: 181
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_7_4
			objectName: "tile_7_4"
            x: 61
            y: 188
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_7_5
			objectName: "tile_7_5"
            x: 62
            y: 180
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_7_6
			objectName: "tile_7_6"
            x: 75
            y: 184
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_7_7
			objectName: "tile_7_7"
            x: 68
            y: 189
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_7_8
			objectName: "tile_7_8"
            x: 76
            y: 189
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_7_9
			objectName: "tile_7_9"
            x: 70
            y: 172
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_7_10
            objectName: "tile_7_10"
            x: 77
            y: 180
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_7_11
            objectName: "tile_7_11"
            x: 68
            y: 187
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_7_12
            objectName: "tile_7_12"
            x: 70
            y: 188
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_7_13
            objectName: "tile_7_13"
            x: 66
            y: 176
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_7_14
            objectName: "tile_7_14"
            x: 67
            y: 172
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_7_15
            objectName: "tile_7_15"
            x: 74
            y: 183
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_7_16
            objectName: "tile_7_16"
            x: 64
            y: 179
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_8_1
			objectName: "tile_8_1"
            x: 65
            y: 173
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_8_2
			objectName: "tile_8_2"
            x: 64
            y: 179
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_8_3
			objectName: "tile_8_3"
            x: 77
            y: 176
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_8_4
			objectName: "tile_8_4"
            x: 78
            y: 186
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_8_5
			objectName: "tile_8_5"
            x: 75
            y: 177
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_8_6
			objectName: "tile_8_6"
            x: 62
            y: 189
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_8_7
			objectName: "tile_8_7"
            x: 70
            y: 186
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_8_8
			objectName: "tile_8_8"
            x: 63
            y: 178
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_8_9
			objectName: "tile_8_9"
            x: 76
            y: 188
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_8_10
            objectName: "tile_8_10"
            x: 68
            y: 180
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_8_11
            objectName: "tile_8_11"
            x: 78
            y: 176
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_8_12
            objectName: "tile_8_12"
            x: 59
            y: 183
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_8_13
            objectName: "tile_8_13"
            x: 61
            y: 182
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_8_14
            objectName: "tile_8_14"
            x: 76
            y: 182
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_8_15
            objectName: "tile_8_15"
            x: 76
            y: 183
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_8_16
            objectName: "tile_8_16"
            x: 72
            y: 173
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_9_1
			objectName: "tile_9_1"
            x: 69
            y: 188
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_9_2
			objectName: "tile_9_2"
            x: 72
            y: 180
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_9_3
			objectName: "tile_9_3"
            x: 65
            y: 184
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_9_4
			objectName: "tile_9_4"
            x: 75
            y: 181
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_9_5
			objectName: "tile_9_5"
            x: 68
            y: 183
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_9_6
			objectName: "tile_9_6"
            x: 73
            y: 190
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_9_7
			objectName: "tile_9_7"
            x: 75
            y: 180
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_9_8
			objectName: "tile_9_8"
            x: 71
            y: 187
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_9_9
			objectName: "tile_9_9"
            x: 73
            y: 175
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_9_10
            objectName: "tile_9_10"
            x: 70
            y: 178
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_9_11
            objectName: "tile_9_11"
            x: 63
            y: 182
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_9_12
            objectName: "tile_9_12"
            x: 63
            y: 184
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_9_13
            objectName: "tile_9_13"
            x: 73
            y: 182
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_9_14
            objectName: "tile_9_14"
            x: 59
            y: 185
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_9_15
            objectName: "tile_9_15"
            x: 66
            y: 176
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_9_16
            objectName: "tile_9_16"
            x: 77
            y: 173
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_10_1
            objectName: "tile_10_1"
            x: 72
            y: 183
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_10_2
            objectName: "tile_10_2"
            x: 70
            y: 176
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_10_3
            objectName: "tile_10_3"
            x: 61
            y: 180
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_10_4
            objectName: "tile_10_4"
            x: 76
            y: 184
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_10_5
            objectName: "tile_10_5"
            x: 70
            y: 177
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_10_6
            objectName: "tile_10_6"
            x: 76
            y: 179
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_10_7
            objectName: "tile_10_7"
            x: 61
            y: 176
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_10_8
            objectName: "tile_10_8"
            x: 73
            y: 184
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_10_9
            objectName: "tile_10_9"
            x: 59
            y: 184
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_10_10
            objectName: "tile_10_10"
            x: 73
            y: 187
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_10_11
            objectName: "tile_10_11"
            x: 64
            y: 184
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_10_12
            objectName: "tile_10_12"
            x: 67
            y: 182
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_10_13
            objectName: "tile_10_13"
            x: 72
            y: 185
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_10_14
            objectName: "tile_10_14"
            x: 72
            y: 181
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_10_15
            objectName: "tile_10_15"
            x: 72
            y: 182
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_10_16
            objectName: "tile_10_16"
            x: 61
            y: 174
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_11_1
            objectName: "tile_11_1"
            x: 67
            y: 189
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_11_2
            objectName: "tile_11_2"
            x: 63
            y: 174
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_11_3
            objectName: "tile_11_3"
            x: 65
            y: 190
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_11_4
            objectName: "tile_11_4"
            x: 66
            y: 187
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_11_5
            objectName: "tile_11_5"
            x: 76
            y: 181
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_11_6
            objectName: "tile_11_6"
            x: 74
            y: 182
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_11_7
            objectName: "tile_11_7"
            x: 65
            y: 190
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_11_8
            objectName: "tile_11_8"
            x: 59
            y: 178
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_11_9
            objectName: "tile_11_9"
            x: 59
            y: 190
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_11_10
            objectName: "tile_11_10"
            x: 59
            y: 172
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_11_11
            objectName: "tile_11_11"
            x: 78
            y: 176
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_11_12
            objectName: "tile_11_12"
            x: 74
            y: 185
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_11_13
            objectName: "tile_11_13"
            x: 77
            y: 184
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_11_14
            objectName: "tile_11_14"
            x: 61
            y: 177
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_11_15
            objectName: "tile_11_15"
            x: 76
            y: 176
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_11_16
            objectName: "tile_11_16"
            x: 68
            y: 175
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_12_1
            objectName: "tile_12_1"
            x: 67
            y: 183
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_12_2
            objectName: "tile_12_2"
            x: 68
            y: 191
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_12_3
            objectName: "tile_12_3"
            x: 71
            y: 174
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_12_4
            objectName: "tile_12_4"
            x: 71
            y: 177
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_12_5
            objectName: "tile_12_5"
            x: 64
            y: 178
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_12_6
            objectName: "tile_12_6"
            x: 77
            y: 188
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_12_7
            objectName: "tile_12_7"
            x: 59
            y: 174
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_12_8
            objectName: "tile_12_8"
            x: 65
            y: 173
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_12_9
            objectName: "tile_12_9"
            x: 60
            y: 190
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_12_10
            objectName: "tile_12_10"
            x: 78
            y: 181
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_12_11
            objectName: "tile_12_11"
            x: 77
            y: 183
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_12_12
            objectName: "tile_12_12"
            x: 72
            y: 181
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_12_13
            objectName: "tile_12_13"
            x: 70
            y: 175
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_12_14
            objectName: "tile_12_14"
            x: 63
            y: 188
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_12_15
            objectName: "tile_12_15"
            x: 59
            y: 185
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_12_16
            objectName: "tile_12_16"
            x: 60
            y: 190
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_13_1
            objectName: "tile_13_1"
            x: 77
            y: 184
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_13_2
            objectName: "tile_13_2"
            x: 75
            y: 193
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_13_3
            objectName: "tile_13_3"
            x: 75
            y: 185
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_13_4
            objectName: "tile_13_4"
            x: 82
            y: 175
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_13_5
            objectName: "tile_13_5"
            x: 85
            y: 175
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_13_6
            objectName: "tile_13_6"
            x: 82
            y: 176
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_13_7
            objectName: "tile_13_7"
            x: 69
            y: 176
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_13_8
            objectName: "tile_13_8"
            x: 71
            y: 180
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_13_9
            objectName: "tile_13_9"
            x: 68
            y: 184
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_13_10
            objectName: "tile_13_10"
            x: 70
            y: 192
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_13_11
            objectName: "tile_13_11"
            x: 79
            y: 178
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_13_12
            objectName: "tile_13_12"
            x: 73
            y: 189
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_13_13
            objectName: "tile_13_13"
            x: 79
            y: 190
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_13_14
            objectName: "tile_13_14"
            x: 68
            y: 190
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_13_15
            objectName: "tile_13_15"
            x: 79
            y: 177
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_13_16
            objectName: "tile_13_16"
            x: 70
            y: 191
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_14_1
            objectName: "tile_14_1"
            x: 80
            y: 190
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_14_2
            objectName: "tile_14_2"
            x: 75
            y: 185
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_14_3
            objectName: "tile_14_3"
            x: 78
            y: 176
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_14_4
            objectName: "tile_14_4"
            x: 86
            y: 175
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_14_5
            objectName: "tile_14_5"
            x: 84
            y: 188
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_14_6
            objectName: "tile_14_6"
            x: 83
            y: 176
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_14_7
            objectName: "tile_14_7"
            x: 71
            y: 175
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_14_8
            objectName: "tile_14_8"
            x: 82
            y: 176
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_14_9
            objectName: "tile_14_9"
            x: 79
            y: 174
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_14_10
            objectName: "tile_14_10"
            x: 82
            y: 192
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_14_11
            objectName: "tile_14_11"
            x: 80
            y: 188
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_14_12
            objectName: "tile_14_12"
            x: 81
            y: 192
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_14_13
            objectName: "tile_14_13"
            x: 82
            y: 177
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_14_14
            objectName: "tile_14_14"
            x: 69
            y: 184
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_14_15
            objectName: "tile_14_15"
            x: 67
            y: 182
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_14_16
            objectName: "tile_14_16"
            x: 71
            y: 188
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_15_1
            objectName: "tile_15_1"
            x: 74
            y: 177
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_15_2
            objectName: "tile_15_2"
            x: 74
            y: 180
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_15_3
            objectName: "tile_15_3"
            x: 78
            y: 180
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_15_4
            objectName: "tile_15_4"
            x: 69
            y: 181
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_15_5
            objectName: "tile_15_5"
            x: 77
            y: 176
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_15_6
            objectName: "tile_15_6"
            x: 80
            y: 186
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_15_7
            objectName: "tile_15_7"
            x: 85
            y: 186
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_15_8
            objectName: "tile_15_8"
            x: 71
            y: 191
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_15_9
            objectName: "tile_15_9"
            x: 70
            y: 177
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_15_10
            objectName: "tile_15_10"
            x: 69
            y: 176
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_15_11
            objectName: "tile_15_11"
            x: 71
            y: 176
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_15_12
            objectName: "tile_15_12"
            x: 86
            y: 192
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_15_13
            objectName: "tile_15_13"
            x: 78
            y: 184
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_15_14
            objectName: "tile_15_14"
            x: 86
            y: 187
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_15_15
            objectName: "tile_15_15"
            x: 79
            y: 185
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_15_16
            objectName: "tile_15_16"
            x: 86
            y: 181
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_16_1
            objectName: "tile_16_1"
            x: 68
            y: 180
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_16_2
            objectName: "tile_16_2"
            x: 78
            y: 193
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_16_3
            objectName: "tile_16_3"
            x: 84
            y: 181
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_16_4
            objectName: "tile_16_4"
            x: 84
            y: 179
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_16_5
            objectName: "tile_16_5"
            x: 85
            y: 174
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_16_6
            objectName: "tile_16_6"
            x: 74
            y: 181
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_16_7
            objectName: "tile_16_7"
            x: 83
            y: 180
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_16_8
            objectName: "tile_16_8"
            x: 73
            y: 190
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_16_9
            objectName: "tile_16_9"
            x: 71
            y: 189
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_16_10
            objectName: "tile_16_10"
            x: 74
            y: 192
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_16_11
            objectName: "tile_16_11"
            x: 74
            y: 185
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_16_12
            objectName: "tile_16_12"
            x: 70
            y: 181
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_16_13
            objectName: "tile_16_13"
            x: 71
            y: 193
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_16_14
            objectName: "tile_16_14"
            x: 76
            y: 187
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_16_15
            objectName: "tile_16_15"
            x: 74
            y: 191
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }

        Image {
            id: image_16_16
			objectName: "tile_16_16"
            x: 74
            y: 194
            width: 32
            height: 32
            source: "image://png/ui/tiles/Tile-000.png"
        }
    }
}
