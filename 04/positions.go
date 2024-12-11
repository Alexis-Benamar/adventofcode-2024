package main

type CharPos struct {
	char string
	x int
	y int
}

var XMASright = []CharPos{{"M", 0, 1}, {"A", 0, 2}, {"S", 0, 3}}
var XMASdown = []CharPos{{"M", 1, 0}, {"A", 2, 0}, {"S", 3, 0}}
var XMASleft = []CharPos{{"M", 0, -1}, {"A", 0, -2}, {"S", 0, -3}}
var XMASup = []CharPos{{"M", -1, 0}, {"A", -2, 0}, {"S", -3, 0}}
var XMASupRight = []CharPos{{"M", -1, 1}, {"A", -2, 2}, {"S", -3, 3}}
var XMASdownRight = []CharPos{{"M", 1, 1}, {"A", 2, 2}, {"S", 3, 3}}
var XMASdownLeft = []CharPos{{"M", 1, -1}, {"A", 2, -2}, {"S", 3, -3}}
var XMASupLeft = []CharPos{{"M", -1, -1}, {"A", -2, -2}, {"S", -3, -3}}

var XMASPositions = [][]CharPos{XMASright, XMASdown, XMASleft, XMASup, XMASupRight, XMASdownRight, XMASdownLeft, XMASupLeft}

var MASdownRight = []CharPos{{"M", -1, -1}, {"S", 1, 1}}
var MASdownLeft = []CharPos{{"M", 1, -1}, {"S", -1, 1}}
var MASupLeft = []CharPos{{"M", 1, 1}, {"S", -1, -1}}
var MASupRight = []CharPos{{"M", -1, 1}, {"S", 1, -1}}

var MASPositions = [][]CharPos{MASdownRight, MASdownLeft, MASupLeft, MASupRight}