package com.example.android_view.page;

import android.os.Bundle;

import androidx.fragment.app.Fragment;

import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Toast;

import com.example.android_view.R;
import com.example.android_view.databinding.FragmentHomePageBinding;
import com.example.genlib.bean.Request;
import com.example.genlib.bean.Response;
import com.example.genlib.service.Service;

import retrofit2.Call;
import retrofit2.Callback;

/**
 * A simple {@link Fragment} subclass.
 * create an instance of this fragment.
 */
public class HomePageFragment extends Fragment {

    private FragmentHomePageBinding binding;

    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
    }

    @Override
    public View onCreateView(LayoutInflater inflater, ViewGroup container,
                             Bundle savedInstanceState) {
        binding = FragmentHomePageBinding.inflate(inflater, container,false);
        binding.callDemoButton.setOnClickListener((e) -> {
            String host = binding.hostEdit.getText().toString();
            String port = binding.portEdit.getText().toString();
            String baseUrl = String.format("http://%s:%s", host, port);
            Service service = new Service(baseUrl);
            Request req = new Request();
            req.setName("you");
            Log.d("[调试]", baseUrl);
            service.demoHandler(req, new Callback<Response>() {
                @Override
                public void onResponse(Call<Response> call, retrofit2.Response<Response> response) {
                    Toast.makeText(getContext(), String.format("onResponse:%s", response.message()), Toast.LENGTH_SHORT).show();
                }

                @Override
                public void onFailure(Call<Response> call, Throwable throwable) {
                    Toast.makeText(getContext(), String.format("onFailure:%s", throwable.getMessage()), Toast.LENGTH_SHORT).show();
                    Log.d("[调试]", throwable.getMessage());
                }
            });
        });
        return binding.getRoot();
    }
}