<?php

namespace App\Http\Controllers\Web;

use App\Http\Controllers\Controller;

class HomeController extends Controller
{

    public function __construct()
    {
    }

    public function index()
    {
        return view('web/home');
    }
}
