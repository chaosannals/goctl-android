package com.example.compose_view.ui.widget

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Text
import androidx.compose.material3.TextField
import androidx.compose.material3.TextFieldDefaults
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp

@Composable
fun LabelTextInput(
    label: String,
    value: String,
    onValueChange:  (String) -> Unit,
) {
    Row(
        horizontalArrangement = Arrangement.Start,
        verticalAlignment = Alignment.CenterVertically,
        modifier = Modifier
            .padding(4.dp)
            .fillMaxWidth()
            .height(64.dp)
            .border(1.dp, Color.Gray, RoundedCornerShape(10.dp))
            .background(Color.White, RoundedCornerShape(10.dp))
            .clip(RoundedCornerShape(10.dp))
    ) {
        Text(
            text = "$label:",
            fontSize = 40.sp,
            modifier = Modifier
                .weight(3f)
        )
        TextField(
            modifier = Modifier
                .weight(7f),
            value = value,
            textStyle = TextStyle(
                fontSize = 32.sp,
            ),
            colors= TextFieldDefaults.colors().copy (
                focusedTextColor= Color.Black,
                unfocusedTextColor=Color.Gray,
                focusedContainerColor=Color.White,
                unfocusedContainerColor = Color.White,
            ),
            onValueChange = onValueChange
        )
    }
}

@Preview
@Composable
fun LabelTextInputPreview() {
    LabelTextInput("host", "192.168.0.1", {})
}