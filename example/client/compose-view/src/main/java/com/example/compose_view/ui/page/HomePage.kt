package com.example.compose_view.ui.page

import android.util.Log
import android.widget.Toast
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.tooling.preview.Preview
import com.example.compose_view.ui.widget.LabelTextInput
import com.example.compose_view.ui.widget.TextButton
import com.example.genlib.bean.Request
import com.example.genlib.bean.Response
import com.example.genlib.service.Service
import retrofit2.Call
import retrofit2.Callback

@Composable
fun HomePage() {
    var host by remember {
        mutableStateOf("192.168.0.1")
    }
    var port by remember {
        mutableStateOf("8888")
    }

    val service = remember(host, port) {
        Service("http://$host:$port")
    }

    Column(
        modifier = Modifier
            .fillMaxSize()
    ) {
        var context = LocalContext.current

        LabelTextInput("host", host) { host = it }
        LabelTextInput("port", port) { port = it }
        TextButton("call demo") {
            val req = Request().apply {
                name = "me"
            }
            service.demoHandler(req, object: Callback<Response> {
                override fun onResponse(
                    p0: Call<Response?>,
                    p1: retrofit2.Response<Response?>
                ) {
                    Toast.makeText(context, "onResponse: ${p1.message()}", Toast.LENGTH_SHORT).show()
                }

                override fun onFailure(
                    p0: Call<Response?>,
                    p1: Throwable
                ) {
                    Toast.makeText(context, "onFailure: ${p1.message}", Toast.LENGTH_SHORT).show()
                    Log.d("[调试]", "onFailure: ${p1.message}")
                }
            })
        }
    }
}

@Preview
@Composable
fun HomePagePreview() {
    HomePage()
}